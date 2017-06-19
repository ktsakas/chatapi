package main

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware jwt authorization middleware
func AuthMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			if (username == "admin" && password == "admin") || (username == "test" && password == "test") {
				return username, true
			}

			return username, false
		},
		Authorizator: func(username string, c *gin.Context) bool {
			if username == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// TokenLookup used to extract the token from the request.
		TokenLookup: "header:Authorization",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc is the current time.
		TimeFunc: time.Now,
	}
}
