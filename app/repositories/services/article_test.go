package services

import (
	"blog-api/app/utils"
	"fmt"
	"testing"
)

func TestGetArticleList(t *testing.T) {
	InitMySQL("../../../app.yaml")

	condition := map[string]string{}
	condition["title"] = "RESTful"

	articles := GetArticleList(condition)
	fmt.Println(utils.JsonMarshalIndent(articles, "	", "	"))

}
