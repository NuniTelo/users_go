package jwt

import (

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}

		token := c.Query("token")
		if token == "" {
			return
		} else {
			_, err := util.ParseToken(token)  //we need to parse the token here
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT  //if false return false
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL //chck this shit out
				}
			}
		}

		if code != e.SUCCESS {  //if success is not
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()   //if succcess then go
	}
}