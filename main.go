package main

import (
	"NBAdata/catch/team"
	"NBAdata/dao"
	_ "NBAdata/inital"
)
func main() {
	//inital 连接db

	//建表
	initTable()

	//抓数据
	//initTeamData()
	//initPlayerData()


	//启动服务
	//api.NewServer().Start()
}

func initTable() {
	//nba球队信息
	//dao.CreateTeam()
	//dao.CreatePlayer()
	dao.CreatePlayerData()
}

func initTeamData() {
	team.PullTeamData("https://nba.hupu.com/teams")
}

func initPlayerData() {
	team.PullPlayerData("https://nba.hupu.com/teams/rockets", 1, "火箭")
}