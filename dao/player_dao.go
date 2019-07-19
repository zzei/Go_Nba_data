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

//获取个人信息
func GetPlayer(player *model.Player) error {

	if err := inital.Db.First(&player).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//按条件查询球员列表
func QueryPlayer(players *[]model.Player, search string, args []interface{}) error {

	if err := inital.Db.Where(search, args[0:]...).Find(&players).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
