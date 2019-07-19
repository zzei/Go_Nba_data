package service

import (
	"NBAdata/dao"
	"NBAdata/model"
)

type SearchPlayerDataParam struct {
	PlayerId int `form:"player_id"`
	Type     int `form:"type"`
}

func QueryPlayerData(param SearchPlayerDataParam) ([]model.PlayerDataSerializer, error) {
	var playerDatas []model.PlayerData

	searchStr := ""
	args := make([]interface{}, 0)

	//查询playerId
	if param.PlayerId >= 0 && param.PlayerId < 1000 {
		searchStr += "player_id = ?"
		args = append(args, param.PlayerId)
	}

	//查询type 1:常规赛 2:季后赛
	if param.Type == 1 || param.Type == 2 {
		if searchStr != "" {
			searchStr += " and "
		}
		searchStr += "(type = ?)"
		args = append(args, param.Type)
	}

	err := dao.QueryPlayerData(&playerDatas, searchStr, args)

	//结果序列化
	var result []model.PlayerDataSerializer
	for _, item := range playerDatas {
		result = append(result, item.Serializer())
	}
	return result, err
}
