package middleware

import (
	"gin-blog/utils"
	"gin-blog/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			code = e.INVALID_PARAMS
		} else {
			token := strings.SplitN(bearerToken, " ", 2)
			claims, err := utils.ParseToken(token[1])
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
