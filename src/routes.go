package main

import (
	"models"

	"github.com/gin-gonic/gin"
)

// SetRoutes unexported
func SetRoutes(r *gin.Engine) {
	db := model.Connect()
	defer db.Close()

	r.PUT("/user/:id", func(c *gin.Context) {
		uuid := c.Param("id")
		var user = model.User{ID: uuid}
		db.First(&user)
	})

	r.GET("/user/:id", func(c *gin.Context) {
		var user model.User
		db.First(&user, c.Param("id"))

		c.JSON(200, user)
	})

	r.GET("/conversation/:user1/:user2", func(c *gin.Context) {
		user1 := c.Query("user1")
		user2 := c.Query("user2")

		var messages = []model.Message{}
		db.Where("sender_uuid = ? AND receiver_uuid = ?", user1, user2).Or("receiver_uuid = ? AND sender_uuid = ?", user1, user2).Find(&messages)

		c.JSON(200, messages)
	})
}
