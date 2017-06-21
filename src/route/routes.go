package route

import (
	"net/http"

	"../chat"
	"../controller"
	"../middleware"
	"../model"

	"github.com/gin-gonic/gin"
)

// GetUserFromClaims returns the user model, given the claims of the authorized user.
func GetUserFromClaims(claims map[string]interface{}) *model.User {
	var email = claims["email"].(string)
	var user, err = model.UserByEmail(email)

	if err != nil {
		// Log critical error here
		// This should never happen
	}

	return user
}

// SetRoutes unexported
func SetRoutes(r *gin.Engine) {
	var authMiddleware = middleware.Auth{
		JwtSecret: []byte("abc"),
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/user", controller.PostUser)

	var auth = r.Group("/")
	// TODO: must validate that the user we need information of is the one logged in
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// The user is authenticated in here
		// however we should still use his credentials to validate he has access
		// to the specific resources

		// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.PUT("/user/:id", controller.PutUser)
		auth.GET("/user/:id", controller.GetUser)

		var hub = chat.New()
		r.GET("/chat", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
			var user = GetUserFromClaims(authMiddleware.GetClaims())
			hub.Serve(user, w, r)
		}))
		// r.POST("/chat/", chat.Handler)
		// r.Handle("WS", "/chat/", chat.Handler)

		auth.GET("/conversation/:channel_id", func(c *gin.Context) {
			// channelId := c.Query("channel_id")

			// var messages = []model.Message{}
			// db.Where("sender_uuid = ? AND receiver_uuid = ?", user1, user2).Or("receiver_uuid = ? AND sender_uuid = ?", user1, user2).Find(&messages)

			// c.JSON(200, messages)
		})
	}
}
