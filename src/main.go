package main

import (
	"./models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := model.Connect()
	defer db.Close()

	r := gin.Default()

	r.PUT("/user", func(c *gin.Context) {
		db.Create(&model.User{})
	})

	r.GET("/user/:id", func(c *gin.Context) {
		var user model.User
		db.First(&user, c.Param("id"))

		c.JSON(200, user)
	})

	r.Run("127.0.0.1:8080")
}
