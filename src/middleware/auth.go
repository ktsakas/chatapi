package middleware

import (
	"fmt"
	"strings"

	"net/http"

	"../model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth is the struct for our authentication middleware
type Auth struct {
	JwtSecret []byte
	claims    map[string]interface{}
}

// SignToken sings the given claims
func (auth *Auth) SignToken(claims jwt.MapClaims) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(auth.JwtSecret)

	return tokenString, err
}

// ValidateToken validates a given token
func (auth *Auth) ValidateToken(tokenString string) (map[string]interface{}, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return auth.JwtSecret, nil
	})

	var claims, ok = token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Authenticator checks the user credentials
// func Authenticator(email, password string) {
// 	var user, loggedIn = model.ValidateUserCredentials(email, password)
// 	if loggedIn {
// 		return user.Email, true
// 	}

// 	return email, false
// }

// LoginHandler handles login requests
// by validating credentials and responding with a signed jwt token
func (auth *Auth) LoginHandler(c *gin.Context) {
	var email = c.PostForm("email")
	var password = c.PostForm("password")
	var user, loggedIn = model.ValidateUserCredentials(email, password)
	if loggedIn {
		var signedToken, err = auth.SignToken(jwt.MapClaims{
			"email": user.Email,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"err":  "internal server error",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"token": signedToken,
			})
		}

		return
	}

	c.JSON(http.StatusForbidden, gin.H{
		"code":    400,
		"message": "invalid email or password",
	})
}

func getBearerToken(c *gin.Context) string {
	return strings.Split(c.Request.Header["Authorization"][0], " ")[1]
}

// GetClaims can only be used after the user has been authenticated
// and returns the claims for the user.
func (auth *Auth) GetClaims() map[string]interface{} {
	if auth.claims == nil {
		return nil
	}

	return auth.claims
}

// MiddlewareFunc gin authorization middleware
func (auth *Auth) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, err = auth.ValidateToken(getBearerToken(c))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err,
			})

			return
		}

		c.Next()
	}

	// return &jwt.GinJWTMiddleware{
	// 	Realm:      "test zone",
	// 	Key:        []byte("secret key"),
	// 	Timeout:    time.Hour,
	// 	MaxRefresh: time.Hour,
	// 	Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
	// 		var user, loggedIn = model.ValidateUserCredentials(username, password)
	// 		if loggedIn {
	// 			return user.Email, true
	// 		}

	// 		return username, false
	// 	},
	// 	Authorizator: func(username string, c *gin.Context) bool {
	// 		var claims = jwt.ExtractClaims(c)
	// 		fmt.Println(claims["uuid"])
	// 		if username != "" {
	// 			return true
	// 		}

	// 		return false
	// 	},
	// 	PayloadFunc: func(userEmail string) map[string]interface{} {
	// 		var userDetails = make(map[string]interface{})
	// 		var user, _ = model.UserByEmail(userEmail)
	// 		userDetails["uuid"] = user.ID
	// 		return userDetails
	// 	},

	// 	// IdentityHandler: func(c *gin.Context) string {
	// 	// 	return "test"
	// 	// },
	// 	Unauthorized: func(c *gin.Context, code int, message string) {
	// 		c.JSON(code, gin.H{
	// 			"code":    code,
	// 			"message": message,
	// 		})
	// 	},

	// 	// TokenLookup used to extract the token from the request.
	// 	TokenLookup: "header:Authorization",

	// 	// TokenHeadName is a string in the header. Default value is "Bearer"
	// 	TokenHeadName: "Bearer",

	// 	// TimeFunc is the current time.
	// 	TimeFunc: time.Now,
	// }
}
