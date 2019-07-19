package team

import (
	"NBAdata/catch"
	"NBAdata/dao"
	"NBAdata/model"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
)

func PullPlayerDataData(url string, playerId int) {
	doc, err := catch.PullData(url)
	if err != nil {
		log.Println(err)
		return
	}

	doc.Find("div.gamecenter_content div.gamecenter_content_l div.team_hig div.shengyasai_tables div.list_table_box").Each(func(i int, selection *goquery.Selection) {
		var dataType int
		if i == 0 {
			dataType = 1
		} else {
			dataType = 2
		}
		selection.Find("table.players_table").Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				//职业生涯数据
				//selection.Find("tr").Eq(2).Find("td").Eq(0).Text()
			} else {
				//赛季数据
				selection.Find("tr").Each(func(j int, selection1 *goquery.Selection) {
					//第一行为标题 略过
					if j != 0 {
						season := selection1.Find("td").Eq(0).Text()
						team := selection1.Find("td").Eq(1).Text()
						timesStr := selection1.Find("td").Eq(2).Text()
						times, _ := strconv.Atoi(timesStr)
						firstTimesStr := selection1.Find("td").Eq(3).Text()
						firstTimes, _ := strconv.Atoi(firstTimesStr)
						floorTimeStr := selection1.Find("td").Eq(4).Text()
						floorTime, _ := strconv.ParseFloat(floorTimeStr, 64)
						shot := selection1.Find("td").Eq(5).Text()
						hit := selection1.Find("td").Eq(6).Text()
						threeShot := selection1.Find("td").Eq(7).Text()
						threeHit := selection1.Find("td").Eq(8).Text()
						penaltyShot := selection1.Find("td").Eq(9).Text()
						penaltyHit := selection1.Find("td").Eq(10).Text()
						reboundStr := selection1.Find("td").Eq(11).Text()
						rebound, _ := strconv.ParseFloat(reboundStr, 64)
						assistStr := selection1.Find("td").Eq(12).Text()
						assist, _ := strconv.ParseFloat(assistStr, 64)
						blockStr := selection1.Find("td").Eq(13).Text()
						block, _ := strconv.ParseFloat(blockStr, 64)
						missStr := selection1.Find("td").Eq(14).Text()
						miss, _ := strconv.ParseFloat(missStr, 64)
						foulStr := selection1.Find("td").Eq(15).Text()
						foul, _ := strconv.ParseFloat(foulStr, 64)
						scoreStr := selection1.Find("td").Eq(16).Text()
						score, _ := strconv.ParseFloat(scoreStr, 64)

						var playerData model.PlayerData
						playerData.PlayerId = playerId
						//type 1为常规赛
						playerData.Type = dataType
						playerData.Season = season
						playerData.Team = team
						playerData.Times = times
						playerData.FirstTimes = firstTimes
						playerData.FloorTime = floorTime
						playerData.Shot = shot
						playerData.Hit = hit
						playerData.ThreeShot = threeShot
						playerData.ThreeHit = threeHit
						playerData.PenaltyShot = penaltyShot
						playerData.PenaltyHit = penaltyHit
						playerData.Rebound = rebound
						playerData.Assist = assist
						playerData.Block = block
						playerData.Miss = miss
						playerData.Foul = foul
						playerData.Score = score

						dao.AddPlayerData(&playerData)
						//fmt.Println(playerData)
					}
				})
			}
		})

	})
}
