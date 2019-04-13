package controllers

import (
	"blog-api/app/repositories/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticleList(c *gin.Context) {
	condition := map[string]string{}
	condition["title"] = c.Query("title")
	condition["content"] = c.Query("content")

	article_list := services.GetArticleList(condition)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功:GetArticleList",
		"data": article_list,
	})
}

func GetArticle(c *gin.Context) {
	article_id, _ := strconv.Atoi(c.Param("id"))
	article := services.GetArticleById(article_id)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功:GetArticleList",
		"data": article,
	})
}
