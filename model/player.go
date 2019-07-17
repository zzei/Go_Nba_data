package model

import "time"

type Player struct {
	Id        int       `gorm:primary_key`
	TeamId    int       `gorm:"type:int;column:team_id"`           //队id
	TeamName  string    `gorm:"type:varchar(50);column:team_name"` //队名
	Number    string    `gorm:"type:varchar(10);column:number"`    //球衣号码
	Name      string    `gorm:"type:varchar(50);column:name"`      //名字
	FaceImg   string    `gorm:"type:varchar(200);column:face_img"` //头像
	Site      string    `gorm:"type:varchar(20);column:site"`      //位置
	Birth     time.Time `gorm:"type:date;column:birth"`            //生日
	CareerAge int       `gorm:"type:int;column:caree_age"`         //职业生涯
	Height    string    `gorm:"type:varchar(50);column:height"`    //身高
	Weight    string    `gorm:"type:varchar(50);column:weight"`    //体重
	Salary    float64   `gorm:"type:double;column:salary"`         //薪资
	School    string    `gorm:"type:varchar(50);column:school"`    //毕业学校
	Draft     string    `gorm:"type:varchar(50);column:draft"`     //选秀
	Country   string    `gorm:"type:varchar(50);column:country"`   //国籍
	Contract  string    `gorm:"type:varchar(200);column:contract"` //合同详情
	CreatedAt time.Time
	DeletedAt *time.Time
}

func (Player) TableName() string {
	return "player"
}

type PlayerSerializer struct {
	Id        int        `json:"id"`
	TeamId    int        `json:"team_id"`
	TeamName  string     `json:"team_name"`
	Number    string     `json:"number"`
	Name      string     `json:"name"`
	FaceImg   string     `json:"face_img"`
	Site      string     `json:"site"`
	Birth     time.Time  `json:"birth"`
	CareerAge int        `json:"career_age"`
	Height    string     `json:"height"`
	Weight    string     `json:"weight"`
	Salary    float64    `json:"salary"`
	School    string     `json:"school"`
	Draft     string     `json:"draft"`
	Country   string     `json:"country"`
	Contract  string     `json:"contract"`
	CreatedAt time.Time  `json:"create_at"`
	DeletedAt *time.Time `json:"delete_at"`
}

func (p *Player) Serializer() PlayerSerializer {
	return PlayerSerializer{
		Id:        p.Id,
		TeamId:    p.TeamId,
		TeamName:  p.TeamName,
		Number:    p.Number,
		Name:      p.Name,
		FaceImg:   p.FaceImg,
		Site:      p.Site,
		Birth:     p.Birth,
		CareerAge: p.CareerAge,
		Height:    p.Height,
		Weight:    p.Weight,
		Salary:    p.Salary,
		School:    p.School,
		Draft:     p.Draft,
		Country:   p.Country,
		Contract:  p.Contract,
		CreatedAt: p.CreatedAt,
		DeletedAt: p.DeletedAt,
	}

}
