package models

import (
)

type Character struct {
	ID 				 uint   `gorm:"primary_key"`
	Name 			 string `gorm:"type: varchar(64); not null; default: ''"`
	Rank 			 string `gorm:"type: varchar(1); not null; default: ''"`
	Element 	 string `gorm:"type: varchar(16); not null; default: ''"`
	Type 			 string `gorm:"type: varchar(16); not null; default: ''"`
	AttackType string `gorm:"type: varchar(16); not null; default: ''"`
	Team			 string `gorm:"type: varchar(16); not null; default: ''"`
	IsLimited  bool 	`gorm:"not null; default: 0"`
}

func (c Character) GetUnitInfo() UnitInfo {
	return UnitInfo{
		"unitType": "character",
		"name": c.Name,
		"rank": c.Rank,
		"element": c.Element,
		"type": c.Type,
		"attackType": c.AttackType,
		"team": c.Team,
	}
}