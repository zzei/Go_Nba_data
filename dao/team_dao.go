package dao

import (
	"NBAdata/inital"
	"NBAdata/model"
	"log"
)

//建表
func CreateTeam() {
	inital.Db.AutoMigrate(&model.Team{})
}

//加队
func AddTeam(team *model.Team) error {
	if err := inital.Db.Create(team).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//获取单队
func GetTeam(team *model.Team) error {

	if err := inital.Db.First(&team).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//获取队伍列表
func QueryTeam(teams *[]model.Team, search string, args []interface{}) error {

	if err := inital.Db.Where(search, args[0:]...).Find(&teams).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

