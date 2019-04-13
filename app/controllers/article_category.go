package controllers

import (
	"blog-api/app/repositories/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategoryList(c *gin.Context) {
	condition := map[string]string{}
	condition["id"] = "1"
	category_list := services.GetCategoryList(condition)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功:GetCategoryList",
		"data": category_list,
	})
}

