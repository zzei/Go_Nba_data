package inital

import (
	"NBAdata/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	dbConf := utils.GetConfig()
	dialect := dbConf.DriverName
	user := dbConf.User
	pwd := dbConf.Pwd
	dbName := dbConf.DbName
	db, err := gorm.Open(dialect, user + ":" + pwd + "@/" + dbName +"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("数据库连接失败")
	}
	db.LogMode(true)
	Db = db
}
