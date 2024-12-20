package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"blog-backend/db"
	"blog-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitPostHandler(client *mongo.Client) {
	// This is now handled in db package
}

func CreatePost(c *gin.Context) {
	log.Printf("Creating new post...")

	// Parse the multipart form with a 32MB max memory
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		log.Printf("Error parsing form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	var post models.Post

	// Get form values
	post.Title = strings.TrimSpace(c.PostForm("title"))
	post.Slug = strings.TrimSpace(c.PostForm("slug"))
	if post.Slug == "" {
		post.Slug = strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))
	}
	post.Summary = strings.TrimSpace(c.PostForm("summary"))
	post.Content = strings.TrimSpace(c.PostForm("content"))
	post.Author = strings.TrimSpace(c.PostForm("author"))

	// Parse read time
	readTime := 5 // Default value
	if rtStr := strings.TrimSpace(c.PostForm("readTime")); rtStr != "" {
		if rt, err := strconv.Atoi(rtStr); err == nil && rt > 0 {
			readTime = rt
		}
	}
	post.ReadTime = readTime

	// Parse published status
	post.Published = c.PostForm("published") == "true"

	// Parse tags
	post.Tags = []string{} // Initialize empty array
	if tagsJSON := strings.TrimSpace(c.PostForm("tags")); tagsJSON != "" {
		var tags []string
		if err := json.Unmarshal([]byte(tagsJSON), &tags); err == nil {
			// Filter out empty tags
			for _, tag := range tags {
				if trimmed := strings.TrimSpace(tag); trimmed != "" {
					post.Tags = append(post.Tags, trimmed)
				}
			}
		}
	}

	// Handle image upload
	if file, err := c.FormFile("image"); err == nil && file != nil {
		if file.Size > 10<<20 { // 10MB limit
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image file too large (max 10MB)"})
			return
		}

		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading image"})
			return
		}
		defer src.Close()

		buffer := make([]byte, file.Size)
		if _, err = src.Read(buffer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading image"})
			return
		}

		post.Image = &models.Image{
			Data:      buffer,
			Filename:  file.Filename,
			MimeType:  file.Header.Get("Content-Type"),
			CreatedAt: time.Now(),
		}
	}

	// Set timestamps
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	// Validate required fields
	if post.Title == "" || post.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and content are required"})
		return
	}

	// Insert into database
	result, err := db.PostsCollection.InsertOne(context.Background(), post)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating post"})
		return
	}

	post.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {
	ctx := context.Background()
	filter := bson.M{}

	// Add published filter for non-admin users
	if !isAdmin(c) {
		filter["published"] = true
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := db.PostsCollection.Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
		return
	}
	defer cursor.Close(ctx)

	var posts []models.Post
	if err := cursor.All(ctx, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	filter := bson.M{"_id": objectID}
	// Add published filter for non-admin users
	if !isAdmin(c) {
		filter["published"] = true
	}

	var post models.Post
	err = db.PostsCollection.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching post"})
		}
		return
	}

	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Parse the multipart form
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		log.Printf("Error parsing form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Get existing post
	var existingPost models.Post
	err = db.PostsCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&existingPost)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Update fields
	post := existingPost
	if title := strings.TrimSpace(c.PostForm("title")); title != "" {
		post.Title = title
	}
	if slug := strings.TrimSpace(c.PostForm("slug")); slug != "" {
		post.Slug = slug
	}
	if summary := strings.TrimSpace(c.PostForm("summary")); summary != "" {
		post.Summary = summary
	}
	if content := strings.TrimSpace(c.PostForm("content")); content != "" {
		post.Content = content
	}
	if author := strings.TrimSpace(c.PostForm("author")); author != "" {
		post.Author = author
	}

	// Update read time
	if rtStr := strings.TrimSpace(c.PostForm("readTime")); rtStr != "" {
		if rt, err := strconv.Atoi(rtStr); err == nil && rt > 0 {
			post.ReadTime = rt
		}
	}

	// Update published status
	post.Published = c.PostForm("published") == "true"

	// Update tags
	if tagsJSON := strings.TrimSpace(c.PostForm("tags")); tagsJSON != "" {
		var tags []string
		if err := json.Unmarshal([]byte(tagsJSON), &tags); err == nil {
			// Filter out empty tags
			post.Tags = []string{}
			for _, tag := range tags {
				if trimmed := strings.TrimSpace(tag); trimmed != "" {
					post.Tags = append(post.Tags, trimmed)
				}
			}
		}
	}

	// Update image if provided
	if file, err := c.FormFile("image"); err == nil && file != nil {
		if file.Size > 10<<20 { // 10MB limit
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image file too large (max 10MB)"})
			return
		}

		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading image"})
			return
		}
		defer src.Close()

		buffer := make([]byte, file.Size)
		if _, err = src.Read(buffer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading image"})
			return
		}

		post.Image = &models.Image{
			Data:      buffer,
			Filename:  file.Filename,
			MimeType:  file.Header.Get("Content-Type"),
			CreatedAt: time.Now(),
		}
	}

	post.UpdatedAt = time.Now()

	// Validate required fields
	if post.Title == "" || post.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and content are required"})
		return
	}

	// Update in database
	_, err = db.PostsCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": post},
	)
	if err != nil {
		log.Printf("Error updating post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	result, err := db.PostsCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting post"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// Helper function to check if user is admin
func isAdmin(c *gin.Context) bool {
	if userValue, exists := c.Get("user_id"); exists {
		if userID, ok := userValue.(string); ok {
			var user models.User
			err := db.UsersCollection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
			return err == nil && user.IsAdmin
		}
	}
	return false
}
