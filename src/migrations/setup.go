package main

import (
	"../models"
)

func main() {
	db := model.Connect()

	// Create model tables
	db.AutoMigrate(&model.User{}, &model.Message{})
}