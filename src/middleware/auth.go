package middleware

import (
	"time"

	"../model"

	"fmt"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// Auth jwt authorization middleware
func Auth() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			var user, loggedIn = model.ValidateUserCredentials(username, password)
			if loggedIn {
				return user.Email, true
			}

			return username, false
		},
		Authorizator: func(username string, c *gin.Context) bool {
			var claims = jwt.ExtractClaims(c)
			fmt.Println(claims["uuid"])
			if username != "" {
				return true
			}

			return false
		},
		PayloadFunc: func(userEmail string) map[string]interface{} {
			var userDetails = make(map[string]interface{})
			var user = model.UserByEmail(userEmail)
			userDetails["uuid"] = user.ID
			return userDetails
		},

		// IdentityHandler: func(c *gin.Context) string {
		// 	return "test"
		// },
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
