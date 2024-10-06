package models

type GachaMachine interface {
	GetInfo() GachaInfo
	Draw(num uint, offsetA uint, offsetS uint, fixedA bool, fixedS bool) []UnitInfo
}

type GachaInfo struct {
	Name string
	PUS_name string
	PUA_name_0 string
	PUA_name_1 string
}
