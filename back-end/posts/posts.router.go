package posts

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine, c *Controller) {
	server.GET("/posts", c.ReadPosts)
	server.POST("/posts", c.CreatePost)
}
