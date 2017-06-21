package controller

import (
	"net/http"

	"../config"
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
		var err = user.Create()
		if err != nil {
			if err == config.ErrRecordExists {
				c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "user already exists"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "bad request"})
	}
}

// PutUser controller updates the user info
func PutUser(c *gin.Context) {
	var id = c.Param("id")
	var user = model.User{ID: id}
	var err = user.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 200, "message": "could not update the user details"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUser controller returns the user info
func GetUser(c *gin.Context) {
	var user, err = model.UserByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
