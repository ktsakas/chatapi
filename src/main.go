package main

import (
	"log"
	"os"
	"time"

	"./model"
	"./route"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

func main() {
	var err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var port = os.Getenv("APP_PORT")
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
	r.Run("127.0.0.1:" + port)
}
