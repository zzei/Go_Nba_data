package api

import (
	"NBAdata/api/team"
	_ "NBAdata/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Api struct {
	engine *gin.Engine
}

func (a *Api) Start() {
	ApiInit(a.engine)
	a.engine.Run(":8088")
}

func ApiInit(r *gin.Engine) {
	//doc

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")

	//team
	teamRegistry(v1)

}

func teamRegistry(c *gin.RouterGroup) {
	c.GET("/team/:id", team.GetTeam)
	c.GET("/teams", team.QueryTeam)

	c.GET("/player/:id", team.GetPlayer)
	c.GET("/players", team.QueryPlayer)

	c.GET("/playerData", team.QueryPlayerData)
}

//启动服务
func NewServer() *Api {
	return &Api{
		engine: gin.Default(),
	}
}
