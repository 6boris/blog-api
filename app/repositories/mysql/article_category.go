package mysql

import "time"

type ArticleCategory struct {
	Id         int        `gorm:"column:id" json:"id"`
	Name       string     `gorm:"column:name" json:"name"`
	Count      int        `gorm:"column:count" json:"count"`
	Created_at *time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (self ArticleCategory) TableName() string {
	return "blog_article_group"
}
