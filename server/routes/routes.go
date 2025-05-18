package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) *gin.Engine {
	// Ping
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "pong",
			})
		})
	}
	return router
}
