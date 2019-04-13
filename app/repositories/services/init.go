package services

import (
	"blog-api/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func InitMySQL(config_path string) {
	c := config.GetConfig(config_path)

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQL.UserName,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Database))

	db = db.LogMode(true)

	//db, err = gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	if err != nil {
		log.Println(err)
	}
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	//return c.MySQL.TablePrefix + defaultTableName
	//	return c.MySQL.TablePrefix + defaultTableName
	//}

	db.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
