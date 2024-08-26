package models

type Unit interface {
	GetUnitInfo()
}

type UnitInfo map[string]string