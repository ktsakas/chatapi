package main

import (
	"./chat"
	"./controller"

	"github.com/gin-gonic/gin"
)

// SetRoutes unexported
func SetRoutes(r *gin.Engine) {

	r.POST("/user", controller.PostUser)
	r.PUT("/user/:id", controller.PutUser)
	r.GET("/user/:id", controller.GetUser)

	var chatServer = chat.New()
	r.GET("/chat/", gin.WrapH(chatServer))
	// r.POST("/chat/", chat.Handler)
	// r.Handle("WS", "/chat/", chat.Handler)

	r.GET("/conversation/:channel_id", func(c *gin.Context) {
		// channelId := c.Query("channel_id")

		// var messages = []model.Message{}
		// db.Where("sender_uuid = ? AND receiver_uuid = ?", user1, user2).Or("receiver_uuid = ? AND sender_uuid = ?", user1, user2).Find(&messages)

		// c.JSON(200, messages)
	})
}
