package services

import (
	"blog-api/app/repositories/mysql"
	"github.com/jinzhu/gorm"
)

func GetArticleList(condition map[string]string) *[]mysql.Article {
	articles := []mysql.Article{}
	db = filterArticle(condition)
	db.Find(&articles)
	return &articles
}

func GetArticleById(id int) *mysql.Article {
	article := mysql.Article{}
	db.Where("id = ?", id).Find(&article)
	return &article
}

//	筛选文章
func filterArticle(condition map[string]string) *gorm.DB {
	if _, ok := condition["title"]; ok {
		db = db.Where("title like ?", "%"+condition["title"]+"%")
	}

	if condition["category_id"] != "" {
		db = db.Where("g_id = ?", condition["g_id"])
	}
	return db
}
