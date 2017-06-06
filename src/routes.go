package main

import (
	"models"

	"github.com/gin-gonic/gin"
)

// SetRoutes unexported
func SetRoutes(r *gin.Engine) {
	db := model.Connect()
	defer db.Close()

	r.PUT("/user", func(c *gin.Context) {
		db.Create(&model.User{})
	})

	r.GET("/user/:id", func(c *gin.Context) {
		var user model.User
		db.First(&user, c.Param("id"))

		c.JSON(200, user)
	})
}
