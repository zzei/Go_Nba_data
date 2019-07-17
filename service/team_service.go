package service

import (
	"NBAdata/dao"
	"NBAdata/model"
	"fmt"
	"strconv"
)


type SearchTeamParam struct {
	Area string		`form:"area"`
	Subarea string	`form:"subarea"`
	Name string		`form:"name"`
}

func GetTeam(id string) (model.TeamSerializer, error){
	var team model.Team
	team.Id, _ = strconv.Atoi(id)

	err := dao.GetTeam(&team)

	return team.Serializer(), err
}

func QueryTeam(param SearchTeamParam) ([]model.TeamSerializer, error) {
	var teams []model.Team

	searchStr := ""
	args := make([]interface{}, 0)

	//查询area
	if param.Area != "" {
		searchStr += "area LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.Area))
	}

	//查询subarea
	if param.Subarea != "" {
		if searchStr != "" {
			searchStr += " and "
		}
		searchStr += "(subarea LIKE ?)"
		args = append(args, fmt.Sprintf("%%%s%%", param.Subarea))
	}

	//查询name
	if param.Name != "" {
		if searchStr != "" {
			searchStr += " and "
		}
		searchStr += "name LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.Name))
	}

	err := dao.QueryTeam(&teams, searchStr, args)

	//结果序列化
	var result []model.TeamSerializer
	for _, item := range teams {
		result = append(result, item.Serializer())
	}
	return result, err
}