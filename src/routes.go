package main

import (
	"net/http"

	"./chat"
	"./controller"
	"./middleware"

	"github.com/gin-gonic/gin"
)

// SetRoutes unexported
func SetRoutes(r *gin.Engine) {
	var authMiddleware = middleware.Auth()

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/user", controller.PostUser)

	var auth = r.Group("/")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.PUT("/user/:id", controller.PutUser)
		auth.GET("/user/:id", controller.GetUser)

		hub := chat.New()
		auth.GET("/chat", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
			hub.Serve(w, r)
		}))
		// r.POST("/chat/", chat.Handler)
		// r.Handle("WS", "/chat/", chat.Handler)

		auth.GET("/conversation/:channel_id", func(c *gin.Context) {
			// channelId := c.Query("channel_id")

			// var messages = []model.Message{}
			// db.Where("sender_uuid = ? AND receiver_uuid = ?", user1, user2).Or("receiver_uuid = ? AND sender_uuid = ?", user1, user2).Find(&messages)

			// c.JSON(200, messages)
		})
	}
}
