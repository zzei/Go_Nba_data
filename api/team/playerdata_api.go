package team

import (
	"NBAdata/service"
	"github.com/gin-gonic/gin"
)

//按条件查询
// @Summary query PlayerData by something
// @Accept json
// @Tags PlayerData
// @Produce  json
// @Param player_id query int true "player_id"
// @Param type query int true "type"
// @Resource PlayerData
// @Router /v1/playerData [get]
// @Success 200 {array} model.PlayerDataSerializer
func QueryPlayerData(c *gin.Context) {

	var param service.SearchPlayerDataParam

	//判断查询参数是否正确
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(400, c.AbortWithError(400, err))
		return
	}

	//调用service
	playerDatas, err := service.QueryPlayerData(param)
	if err != nil {
		c.JSON(400, c.AbortWithError(400, err))
		return
	}

	c.JSON(200, playerDatas)
}
