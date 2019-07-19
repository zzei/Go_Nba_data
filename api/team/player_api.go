package team

import (
	"NBAdata/service"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"log"
)

//getPlayerById
//按条件查询
// @Summary get Player by id
// @Accept json
// @Tags Player
// @Produce  json
// @Param id path int true "id"
// @Resource Player
// @Router /v1/player/{id} [get]
// @Success 200 {object} model.PlayerSerializer
func GetPlayer(c *gin.Context) {
	id := c.Param("id")

	//验证数据正确性
	if id == "" {
		c.JSON(400, c.AbortWithError(400, errors.New("请输入id！")))
		return
	}

	player, err := service.GetPlayer(id)
	if err != nil {
		c.JSON(400, c.AbortWithError(400, errors.New("没有数据！")))
		return
	}

	c.JSON(400, player)
}

//按条件查询
// @Summary query Player by something
// @Accept json
// @Tags Player
// @Produce  json
// @Param name query string false "name"
// @Param team_name query string false "team_name"
// @Param site query string false "site"
// @Resource Player
// @Router /v1/players [get]
// @Success 200 {array} model.PlayerSerializer
func QueryPlayer(c *gin.Context) {
	var param service.SearchPlayerParam

	//判断查询参数是否正确
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(400, c.AbortWithError(400, err))
		return
	}
	log.Println("param", param)

	//调用service
	players, err := service.QueryPlayer(param)
	if err != nil {
		c.JSON(400, c.AbortWithError(400, err))
		return
	}

	c.JSON(200, players)
}
