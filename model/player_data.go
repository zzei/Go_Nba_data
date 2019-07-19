package model

import "time"

type PlayerData struct {
	Id          int        `gorm"type:int(11);column:id"`
	PlayerId    int        `gorm"type:int(11);column:player_id"`
	Season      string     `gorm"type:varchar(20);column:season"`
	Team        string     `gorm"type:varchar(20);column:team"`
	Times       int        `gorm"type:int(11);column:times"`
	FirstTimes  int        `gorm"type:int(11);column:first_times"`
	FloorTime   float64    `gorm"type:double;column:floor_time"`
	Shot        string     `gorm"type:varchar(20);column:shot"`
	Hit         string     `gorm"type:varchar(20);column:hit"`
	ThreeShot   string     `gorm"type:varchar(20);column:three_shot"`
	ThreeHit    string     `gorm"type:varchar(20);column:three_hit"`
	PenaltyShot string     `gorm"type:varchar(20);column:penalty_shot"`
	PenaltyHit  string     `gorm"type:varchar(20);column:penalty_hit"`
	Rebound     float64    `gorm"type:double;column:rebound"`
	Assist      float64    `gorm"type:double;column:assist"`
	Steal       float64    `gorm"type:double;column:steal"`
	Block       float64    `gorm"type:double;column:block"`
	Miss        float64    `gorm"type:double;column:miss"`
	Foul        float64    `gorm"type:double;column:foul"`
	Score       float64    `gorm"type:double;column:score"`
	Type        int        `gorm"type:int(11);column:type"`
	CreatedAt   time.Time  `gorm"type:timestamp;column:created_at"`
	DeletedAt   *time.Time `gorm"type:timestamp;column:deleted_at"`
}

type PlayerDataSerializer struct {
	Id          int        `json"id"`
	PlayerId    int        `json"player_id"`
	Season      string     `json"season"`
	Team        string     `json"team"`
	Times       int        `json"times"`
	FirstTimes  int        `json"first_times"`
	FloorTime   float64    `json"floor_time"`
	Shot        string     `json"shot"`
	Hit         string     `json"hit"`
	ThreeShot   string     `json"three_shot"`
	ThreeHit    string     `json"three_hit"`
	PenaltyShot string     `json"penalty_shot"`
	PenaltyHit  string     `json"penalty_hit"`
	Rebound     float64    `json"rebound"`
	Assist      float64    `json"assist"`
	Steal       float64    `json"steal"`
	Block       float64    `json"block"`
	Miss        float64    `json"miss"`
	Foul        float64    `json"foul"`
	Score       float64    `json"score"`
	Type        int        `json"type"`
	CreatedAt   time.Time  `json"created_at"`
	DeletedAt   *time.Time `json"deleted_at"`
}

func (t *PlayerData) Serializer() PlayerDataSerializer {
	return PlayerDataSerializer{
		Id:          t.Id,
		PlayerId:    t.PlayerId,
		Season:      t.Season,
		Team:        t.Team,
		Times:       t.Times,
		FirstTimes:  t.FirstTimes,
		FloorTime:   t.FloorTime,
		Shot:        t.Shot,
		Hit:         t.Hit,
		ThreeShot:   t.ThreeShot,
		ThreeHit:    t.ThreeHit,
		PenaltyShot: t.PenaltyShot,
		PenaltyHit:  t.PenaltyHit,
		Rebound:     t.Rebound,
		Assist:      t.Assist,
		Steal:       t.Steal,
		Block:       t.Block,
		Miss:        t.Miss,
		Foul:        t.Foul,
		Score:       t.Score,
		Type:        t.Type,
		CreatedAt:   t.CreatedAt,
		DeletedAt:   t.DeletedAt,
	}
}
