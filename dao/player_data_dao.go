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
