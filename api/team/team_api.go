package team

import (
	"NBAdata/service"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"log"
)



//getTeamById
//按条件查询
// @Summary get Team by id
// @Accept json
// @Tags Team
// @Produce  json
// @Param id path int true "id"
// @Resource Team
// @Router /v1/team/{id} [get]
// @Success 200 {object} model.TeamSerializer
func GetTeam(c *gin.Context) {
	id := c.Param("id")

	//验证数据正确性
	if id == "" {
		c.JSON(400, c.AbortWithError(400, errors.New("请输入id！")))
		return
	}

	team, err := service.GetTeam(id)
	if err != nil  {
		c.JSON(400, c.AbortWithError(400, errors.New("没有数据！")))
		return
	}

	c.JSON(400, team)
}


//按条件查询
// @Summary query Team by something
// @Accept json
// @Tags Team
// @Produce  json
// @Param area query string false "area"
// @Param subarea query string false "subarea"
// @Param name query string false "name"
// @Resource Team
// @Router /v1/teams [get]
// @Success 200 {array} model.TeamSerializer
func QueryTeam(c *gin.Context) {
	var param service.SearchTeamParam

	//判断查询参数是否正确
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(400, c.AbortWithError(400, err))
		return
	}
	log.Println("param",param)

	//调用service
	teams, err := service.QueryTeam(param)
	if err != nil {
		c.JSON(400, c.AbortWithError(400, err))
		return
	}

	c.JSON(200, teams)
}