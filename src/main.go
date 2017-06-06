package main

import (
	"github.com/gin-gonic/gin"
	"./models"
)

func main() {
	db := model.Connect()
	defer db.Close()
	
	r := gin.Default()
	r.PUT("/user", func(c *gin.Context) {
		db.Create(&model.User{})
	})

	r.Run("127.0.0.1:8080")
}