package routes

import (
	"net/http"

	"bluebell/controllers"
	"bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)

	api := r.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())
	{
		api.GET("/ping", func(c *gin.Context) {
			userID, _ := c.Get("userID")

			c.JSON(http.StatusOK, gin.H{
				"msg":     "pong",
				"user_id": userID, // 看看能不能拿到
			})
		})
		api.GET("/community", controllers.CommunityHandler)
		api.GET("/community/:id", controllers.CommunityDetailHandler)
	}
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
