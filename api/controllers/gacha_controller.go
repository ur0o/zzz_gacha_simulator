package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"zzz_gacha/models"
)

func IndexGacha(c *gin.Context) {
	info := models.GetGachaMachineAll()
	var res []map[string]interface{}

	for _, i := range info {
		res = append(res, map[string]interface{}{
			"id": i.ID,
			"name": i.Name,
			"puS": i.PUS_name,
			"puA_name_0": i.PUA_name_0,
			"puA_name_1": i.PUA_name_1,
			"startDate": i.StartDate,
			"endDate": i.EndDate,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"results": res,
	})
}

func GetGacha(c *gin.Context) {
	gacha_id, _ := strconv.Atoi(c.Param("id"))

	gacha_machine := models.GetCharacterGacha(gacha_id)
	info := gacha_machine.GetInfo()
	c.JSON(http.StatusOK, gin.H{
		"name": info.Name,
		"puS": info.PUS_name,
		"puA_name_0": info.PUA_name_0,
		"puA_name_1": info.PUA_name_1,
	})
}

func DrawGacha(c *gin.Context) {
	// gacha_type := c.Param("type")
	gacha_id, _ := strconv.Atoi(c.Param("id"))
	oS, _ := strconv.Atoi(c.Param("offsetS"))
	oA, _ := strconv.Atoi(c.Param("offsetA"))
	fS, _ := strconv.ParseBool(c.Param("fixedS"))
	fA, _ := strconv.ParseBool(c.Param("fixedA"))

	var gacha_machine models.GachaMachine
	if true {
		gacha_machine = models.GetCharacterGacha(gacha_id)
	// } else {
	// 	gacha_machine = models.GetWeaponGacha(gacha_id)
	}
	c.JSON(http.StatusOK, gacha_machine.Draw(10, uint(oS), uint(oA), fS, fA))
}
