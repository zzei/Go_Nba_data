package dao

import (
	"NBAdata/inital"
	"NBAdata/model"
	"log"
)

//建表
func CreatePlayerData() {
	inital.Db.AutoMigrate(&model.PlayerData{})
}

//加队
func AddPlayerData(playerData *model.PlayerData) error {
	if err := inital.Db.Create(playerData).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//按条件获取球员数据
func QueryPlayerData(playerDatas *[]model.PlayerData, search string, args []interface{}) error {

	if err := inital.Db.Where(search, args[0:]...).Find(&playerDatas).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
