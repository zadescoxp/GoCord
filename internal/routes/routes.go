package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{

			v1.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Welcome to GoCord🔥",
				})
			})

			v1.GET("/health", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "System is healthy",
				})
			})

			v1.GET("/info", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"name":    "GoCord",
					"version": "1.0.0",
					"author":  "Your Name",
				})
			})
		}
	}

	return router
}
