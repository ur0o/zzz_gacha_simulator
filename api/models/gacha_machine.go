package models

import (
	"zzz_gacha/database"
)

type GachaMachine interface {
	GetInfo() GachaInfo
	Draw(num uint, offsetA uint, offsetS uint, fixedA bool, fixedS bool) []UnitInfo
}

type GachaInfo struct {
	ID uint
	Name string
	PUS_name string
	PUA_name_0 string
	PUA_name_1 string
	StartDate string
	EndDate string
}

func GetGachaMachineAll() []GachaInfo {
	var cg []CharacterGacha
	db := database.GetDB()
	db.Joins("PUS").Joins("PUA_0").Joins("PUA_1").Find(&cg)

	var cgi []GachaInfo
	for _, g := range cg {
		cgi = append(cgi, g.GetInfo())
	}
	return cgi
}