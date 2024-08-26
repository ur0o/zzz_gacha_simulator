package models

type Weapon struct {
	ID 				uint 	 `gorm:"primary_key"`
	Name 			string `gorm:"type: varchar(64); not null; default: ''"`
	Rank 			string `gorm:"type: varchar(1); not null; default: ''"`
	Type 			string `gorm:"type: varchar(16); not null; default: ''"`
	IsLimited bool 	 `gorm:"not null; default: 0;"`
}

func (w Weapon) GetUnitInfo() UnitInfo {
	return UnitInfo{
		"unitType": "weapon",
		"name": w.Name,
		"rank": w.Rank,
		"type": w.Type,
	}
}