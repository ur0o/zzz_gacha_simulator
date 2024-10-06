package models

type Unit interface {
	GetUnitInfo() UnitInfo
}

type UnitInfo map[string]string