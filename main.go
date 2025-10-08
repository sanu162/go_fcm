package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// Initialize Firebase App
func setupFirebase() *firebase.App {
	opt := option.WithCredentialsFile("serviceAccountKey.json") // Path to your Firebase service account key
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}

func main() {
	r := gin.Default()
	firebaseApp := setupFirebase()

	r.POST("/send", func(c *gin.Context) {
		var req struct {
			Token string `json:"token"`
			Title string `json:"title"`
			Body  string `json:"body"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		ctx := context.Background()
		client, err := firebaseApp.Messaging(ctx)
		if err != nil {
			log.Printf("error getting messaging client: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get FCM client"})
			return
		}

		// Build notification
		message := &messaging.Message{
			Token: req.Token,
			Notification: &messaging.Notification{
				Title: req.Title,
				Body:  req.Body,
			},
		}

		// Send notification
		response, err := client.Send(ctx, message)
		if err != nil {
			log.Printf("error sending message: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
			return
		}

		fmt.Printf("Successfully sent message: %s\n", response)
		c.JSON(http.StatusOK, gin.H{"success": true, "response": response})
	})

	r.Run(":8080") // Run server on localhost:8080
}
