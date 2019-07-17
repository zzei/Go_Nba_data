package team

import (
	"NBAdata/catch"
	"NBAdata/dao"
	"NBAdata/model"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

//抓取队数据
func PullTeamData(url string) {
	doc, err := catch.PullData(url)
	if err != nil {
		log.Println(err)
		return
	}

	doc.Find("div.gamecenter_livestart div.gamecenter_content div.teamlist_box").Each(func(i int, selection *goquery.Selection) {
		//西部
		area := selection.Find("div.teamlist_box_l").Text()

		selection.Find("div.teamlist_box_r div.all").Each(func(j int, selection1 *goquery.Selection) {
			subarea := selection1.Find("div.a span").Text()

			selection1.Find("div .team a.a_teamlink").Each(func(k int, selection2 *goquery.Selection) {
				logo,_ := selection2.Find("div.box div.img img").Attr("src")
				name := selection2.Find("div.box div.font h2").Text()

				//抓取详情
				href,_ := selection2.Attr("href")
				doc1, _ := catch.PullData(href)
				allName := doc1.Find("div.gamecenter_content div.team_data h2 span").Eq(0).Text()

				stadium := doc1.Find("div.gamecenter_content div.team_data div.content div.content_a div.font p").Eq(1).Text()
				stadium = stadium[strings.Index(stadium, "：")+3 : strings.Index(stadium, "分区")]

				website := doc1.Find("div.gamecenter_content div.team_data div.content div.content_a div.font p a").Text()
				info := doc1.Find("div.gamecenter_content div.team_data div.content div.content_a div.txt").Text()

				team := model.Team{}
				team.Area = area
				team.Subarea = subarea
				team.Logo = logo
				team.Name = name
				team.AllName = allName
				team.Stadium = stadium
				team.Website = website
				team.Info = info

				dao.AddTeam(&team)
				//添加球队后 开始添加该队下球员
				PullPlayerData(href, team.Id, name)
			})
		})

	})
}
