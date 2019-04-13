package routes

import (
	"blog-api/app/controllers"
	"github.com/gin-gonic/gin"
)

func InitIndexRouters(app *gin.Engine) {
	app.GET("index/articles", controllers.GetArticleList)
	app.GET("index/articles/:id", controllers.GetArticle)
	app.GET("index/article_category", controllers.GetCategoryList)
}
