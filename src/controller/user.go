package controller

import (
	"net/http"

	"../model"

	"github.com/gin-gonic/gin"
)

type signupForm struct {
	Email     string `form:"email" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Sex       string `form:"sex" binding:"required"`
	TalkingTo string `form:"talkingTo" binding:"required"`
}

// PostUser controller creates a new user
func PostUser(c *gin.Context) {
	var form signupForm

	if c.Bind(&form) == nil {
		var user = model.User{
			Email:     form.Email,
			Password:  form.Password,
			Sex:       form.Sex,
			TalkingTo: form.TalkingTo,
		}
		user.Create()

		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
	}
}

// PutUser controller updates the user info
func PutUser(c *gin.Context) {
	uuid := c.Param("id")
	var user = model.User{ID: uuid}
	user.Create()
}

// GetUser controller returns the user info
func GetUser(c *gin.Context) {
	var user = model.UserByID(c.Param("id"))

	c.JSON(200, user)
}
