package services

import "blog-api/app/repositories/mysql"

func GetCategoryList(condition map[string]string) *[]mysql.ArticleCategory {
	category_list := []mysql.ArticleCategory{}

	db.Table("blog_article_group").
		Select("" +
			"blog_article_group.id, " +
			"blog_article_group.name," +
			"count(blog_article.id) AS count, " +
			"blog_article_group.created_at, " +
			"blog_article_group.updated_at").
		Joins("left join blog_article on blog_article_group.id = blog_article.g_id").
		Group("blog_article_group.id").
		Find(&category_list)

	return &category_list
}
