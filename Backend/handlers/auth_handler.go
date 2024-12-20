package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"blog-backend/db"
	"blog-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func InitAuthHandler(client *mongo.Client) {
	// Remove this since it's handled in db package
}

func HandleLogin(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		log.Printf("Invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("Login attempt for email: %s", loginData.Email)

	// Find user
	var user models.User
	err := db.UsersCollection.FindOne(context.Background(), bson.M{"email": loginData.Email}).Decode(&user)
	if err != nil {
		log.Printf("User not found: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		log.Printf("Password mismatch: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	log.Printf("Login successful for user: %s", user.Email)

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"email":   user.Email,
		"isAdmin": user.IsAdmin,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":      user.ID,
			"email":   user.Email,
			"isAdmin": user.IsAdmin,
		},
	})
}
