package main

import (
	"blog-backend/db"
	"blog-backend/handlers"
	"blog-backend/middleware"
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI is not set in environment variables")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	log.Println("Connected to MongoDB!")

	// Initialize collections
	db.InitCollections(client)

	// Initialize router
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	api := r.Group("/api")
	{
		// Public routes
		api.POST("/contact", handlers.HandleContact)
		api.POST("/login", handlers.HandleLogin)
		api.GET("/posts", handlers.GetPosts)
		api.GET("/posts/:id", handlers.GetPost)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/posts", handlers.CreatePost)
			protected.PUT("/posts/:id", handlers.UpdatePost)
			protected.DELETE("/posts/:id", handlers.DeletePost)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
