package dao

import (
	"NBAdata/inital"
	"NBAdata/model"
	"log"
)

//建表
func CreatePlayer() {
	inital.Db.AutoMigrate(&model.Player{})
}

//加队
func AddPlayer(player *model.Player) error {
	if err := inital.Db.Create(player).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
