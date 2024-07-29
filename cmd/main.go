package main

import (
	database "chatting/pkg/Database"
	"chatting/pkg/api/controllers"
	"chatting/pkg/api/handler"
	"chatting/pkg/repositories"
	"chatting/pkg/usecases/message"
	"chatting/pkg/usecases/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dbClient := database.Connect()

	userRepo := &repositories.UserRepository{Collection: dbClient.Database("chatapp").Collection("users")}
	messageRepo := &repositories.MessageRepository{Collection: dbClient.Database("chatapp").Collection("messages")}

	userInteractor := user.UserInteractor{UserRepository: *userRepo}
	messageInteractor := message.MessageInteractor{MessageRepository: *messageRepo}

	userController := controllers.UserController{UserInteractor: userInteractor}
	messageController := handler.MessageController{MessageInteractor: messageInteractor}

	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUser)
	router.POST("/messages", messageController.SendMessage)
	router.GET("/messages/:senderID/:receiverID", messageController.GetMessages)

	// Serve static files
	router.Static("/static", "./web/static")

	// Load HTML templates
	router.LoadHTMLGlob("web/templates/*")

	router.GET("/sender", func(c *gin.Context) {
		c.HTML(200, "sender.html", nil)
	})

	router.GET("/receiver", func(c *gin.Context) {
		c.HTML(200, "receiver.html", nil)
	})

	router.Run(":8080")
}
