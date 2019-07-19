# Go_Nba_data
NBA's data from hupu
从虎扑抓取的NBA数据
球队 team
球员 player
球员数据 playerData

主体框架采用gin
orm框架采用gorm
爬虫使用goquery
接口文档用swaggor生成


接口：

  /v1/team/{id}      get Team by id
  /v1/teams          query Team by something
  
  /v1/player/{id}    get Player by id
  /v1/players        query Player by something
  
  /v1/playerData     query PlayerData by something
