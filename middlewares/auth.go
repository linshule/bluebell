package middlewares

import (
	"net/http"
	"strings"

	"bluebell/controllers"
	"bluebell/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, http.StatusUnauthorized, "请求头中auth为空")
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, http.StatusUnauthorized, "请求头中auth格式有误")
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, http.StatusUnauthorized, "无效的Token")
			c.Abort()
			return
		}

		c.Set("userID", mc.UserID)

		c.Next()
	}
}
