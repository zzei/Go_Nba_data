package model

import (
	"time"
)

type Team struct {
	Id        int    `gorm:primary_key`
	City      string `gorm:"type:varchar(50);column:city"`                //城市
	Name      string `gorm:"type:varchar(50);not null;column:name"`       //队名
	AllName   string `gorm:"type:varchar(100);not null;column:all_name""` //全名
	Area      string `gorm:"type:varchar(20);not null;column:area"`       //东西分区
	Subarea   string `gorm:"type:varchar(20);not null;column:subarea"`    //分区
	Logo      string `gorm:"type:varchar(200);not null;column:logo"`      //logo图标链接
	Stadium   string `gorm:"type:varchar(50);not null;column:stadium"`    //主场
	Website   string `gorm:"type:varchar(100);column:website"`            //官网
	Info      string `gorm:"type:varchar(200);column:info"`               //简介
	CreatedAt time.Time
	DeletedAt *time.Time
}

func (Team) TableName() string {
	return "team"
}

type TeamSerializer struct {
	Id        int        `json:"id"`
	City      string     `json:"city"`
	Name      string     `json:"name"`
	AllName   string     `json:"all_name"`
	Area      string     `json:"area"`
	Subarea   string     `json:"subarea"`
	Logo      string     `json:"logo"`
	Stadium   string     `json:"stadium"`
	Website   string     `json:"website"`
	Info      string     `json:"info"`
	CreatedAt time.Time  `json:"create_at"`
	DeletedAt *time.Time `json:"delete_at"`
}

func (t *Team) Serializer() TeamSerializer {
	return TeamSerializer{
		Id:        t.Id,
		City:      t.City,
		Name:      t.Name,
		AllName:   t.AllName,
		Area:      t.Area,
		Subarea:   t.Subarea,
		Logo:      t.Logo,
		Stadium:   t.Stadium,
		Website:   t.Website,
		Info:      t.Info,
		CreatedAt: t.CreatedAt,
		DeletedAt: t.DeletedAt,
	}
}
