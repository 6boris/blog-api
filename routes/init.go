package routes

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gin-contrib/cors.v1"
)

func InitRouter(app *gin.Engine) *gin.Engine {
	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	app.Use(cors.Default())
	InitIndexRouters(app)
	return app
}
