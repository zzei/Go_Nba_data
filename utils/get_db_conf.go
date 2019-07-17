package utils

import (
	"encoding/json"
	"log"
	"os"
)

type MysqlDb struct {
	DriverName string
	User string
	Pwd string
	Host string
	Port string
	DbName string
}


//读取配置文件 返回
func GetConfig() MysqlDb{

	path := "conf/db.json"
	file, _ := os.Open(path)
	defer file.Close()

	var conf MysqlDb
	deCoder := json.NewDecoder(file)
	err := deCoder.Decode(&conf)
	if err != nil {
		log.Println(err)
	}
	return conf
}
