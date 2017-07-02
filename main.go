package main

import (
	"time"

	"./config"
	"./model"
	"./route"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

func main() {
	config.LoadConfig()
	model.Connect("collegechat")

	var r = gin.Default()

	// Allow cross-origin
	// TODO: only in test mode
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	route.SetRoutes(r)
	r.Run("127.0.0.1:" + config.DevPort)
}
