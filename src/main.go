package main

import (
	"./model"
	"github.com/gin-gonic/gin"
)

func main() {
	model.Connect()

	r := gin.Default()
	SetRoutes(r)
	r.Run("127.0.0.1:8080")
}
