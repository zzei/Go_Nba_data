package service

import (
	"NBAdata/dao"
	"NBAdata/model"
	"fmt"
	"strconv"
)

type SearchPlayerParam struct {
	Name     string `form:"name"`
	TeamName string `form:"team_name"`
	Site     string `form:"site"`
}

func GetPlayer(id string) (model.PlayerSerializer, error) {
	var player model.Player
	player.Id, _ = strconv.Atoi(id)

	err := dao.GetPlayer(&player)

	return player.Serializer(), err
}

func QueryPlayer(param SearchPlayerParam) ([]model.PlayerSerializer, error) {
	var players []model.Player

	searchStr := ""
	args := make([]interface{}, 0)

	//查询area
	if param.Name != "" {
		searchStr += "name LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.Name))
	}

	//查询subarea
	if param.TeamName != "" {
		if searchStr != "" {
			searchStr += " and "
		}
		searchStr += "(team_name = ?)"
		args = append(args, param.TeamName)
	}

	//查询name
	if param.Site != "" {
		if searchStr != "" {
			searchStr += " and "
		}
		searchStr += "site LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.Site))
	}

	err := dao.QueryPlayer(&players, searchStr, args)

	//结果序列化
	var result []model.PlayerSerializer
	for _, item := range players {
		result = append(result, item.Serializer())
	}
	return result, err
}
