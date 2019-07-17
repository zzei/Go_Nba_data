package team

import (
	"NBAdata/catch"
	"NBAdata/dao"
	"NBAdata/model"
	"NBAdata/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
	"time"
)

func PullPlayerData(url string, teamId int, teamName string) {
	doc, err := catch.PullData(url)
	if err != nil {
		log.Println(err)
		return
	}

	//抓取球员个人链接
	doc.Find("div.gamecenter_content div.team_list_data div.jiben_title_table div.a div").Each(func(i int, selection *goquery.Selection) {
		//第一行为title跳过
		if i != 0 {
			number := selection.Find("span.c1").Text()

			href, _ := selection.Find("span.c2 a").Attr("href")
			doc1, _ := catch.PullData(href)
			name := doc1.Find("div.gamecenter_content div.team_data h2").Text()
			name = utils.FilterEnter(name[0:strings.Index(name,"）")+3])
			img, _ := doc1.Find("div.gamecenter_content div.team_data div.content div.content_a div.img img").Attr("src")
			var site string
			var height string
			var weight string
			var birth time.Time
			var school string
			var draft string
			var country string
			var salary float64
			var contract string

			doc1.Find("div.gamecenter_content div.team_data div.content div.content_a div.font p").Each(func(j int, selection1 *goquery.Selection) {
				str := selection1.Text()
				if strings.Contains(str, "位置") {
					site = str[strings.Index(str,"：")+3:]
				} else if strings.Contains(str, "身高") {
					height = str[strings.Index(str,"：")+3:]
				} else if strings.Contains(str, "体重") {
					weight = str[strings.Index(str,"：")+3:]
				} else if strings.Contains(str, "生日") {
					birthStr := str[strings.Index(str, "：")+3:]
					birth,_ = time.Parse("2006-01-02", birthStr)
				} else if strings.Contains(str, "学校") {
					school = str[strings.Index(str,"：")+3:]
				} else if strings.Contains(str, "选秀") {
					draft = str[strings.Index(str,"：")+3:]
				} else if strings.Contains(str, "国籍") {
					country = str[strings.Index(str,"：")+3:]
				} else if strings.Contains(str, "本赛季薪金") {
					salaryStr := str[strings.Index(str,"：")+3:strings.Index(str, "万")]
					salary,_ = strconv.ParseFloat(salaryStr, 64)
				} else if strings.Contains(str, "合同") {
					contract = str[strings.Index(str,"：")+3:]
				} else {

				}
			})

			var player model.Player

			player.TeamId = teamId
			player.TeamName = teamName
			player.Number = number
			player.Name = name
			player.FaceImg = img
			player.Site = site
			player.Height = height
			player.Weight = weight
			player.Birth = birth
			player.School = school
			player.Draft = draft
			player.Country = country
			player.Salary = salary
			player.Contract = contract

			//fmt.Println(player.Id)

			dao.AddPlayer(&player)
		}
	})
}
