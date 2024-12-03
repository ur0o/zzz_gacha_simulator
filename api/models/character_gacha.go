package models

import (
	"math"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"zzz_gacha/database"
)

type CharacterGacha struct {
	gorm.Model
	ID 				uint 			`gorm:"primary_key"`
	Name 			string 		`gorm:"type:varchar(64); not null; default: ''"`
	PUS_ID 		uint
	PUS 			Character
	PUA_0ID 	uint
	PUA_0 		Character
	PUA_1ID 	uint
	PUA_1 		Character
	StartDate time.Time `gorm:"type:date; not null;"`
	EndDate 	time.Time `gorm:"type:date; not null;"`
}

func GetCharacterGacha(id int) GachaMachine {
	var gacha CharacterGacha
	db := database.GetDB()

	db.Where(&CharacterGacha{ID: uint(id)}).Find(&gacha)
	return gacha
}

func (g CharacterGacha)GetInfo() GachaInfo {
	return GachaInfo{
		ID: g.ID,
		Name: g.Name,
		PUS_name: g.PUS.Name,
		PUA_name_0: g.PUA_0.Name,
		PUA_name_1: g.PUA_1.Name,
		StartDate: g.StartDate.Format("2006/01/02"),
		EndDate: g.EndDate.Format("2006/01/02"),
	}
}

func (g CharacterGacha)Draw(n uint, oS uint, oA uint, fS bool, fA bool) []UnitInfo {
	// setup
	var puAcs []Character
	var Scs []Character
	var Acs []Character
	var Aws	[]Weapon
	var Bws []Weapon

	db := database.GetDB()
	db.Preload("PUS").Find(&CharacterGacha{})
	db.Where([]uint{g.PUA_0ID, g.PUA_1ID}).Find(&puAcs)
	db.Where(&Character{Rank: "S", IsLimited: false}).Not(g.PUS_ID).Find(&Scs)
	db.Where(&Character{Rank: "A", IsLimited: false}).Not([]uint{g.PUA_0ID, g.PUA_1ID}).Find(&Acs)
	db.Where(&Weapon{Rank: "A", IsLimited: false}).Find(&Aws)
	db.Where(&Weapon{Rank: "B", IsLimited: false}).Find(&Bws)

	// gacha
	coS := oS
	coA := oA
	cfS := fS
	cfA := fA
	var results []UnitInfo
	for i := 0; i < int(n); i++ {
		rank := drawRank(coA, coS)
		if rank == "S" {
			coS = 0
			coA = 0
			if cfS || drawPU() {
				cfS = false
				results = append(results, g.PUS.GetUnitInfo())
			} else {
				cfS = true
				results = append(results, Scs[rand.Intn(len(Scs))].GetUnitInfo())
			}
		} else if rank == "A" {
			coS++
			coA = 0
			if cfA || drawPU() {
				cfA = false
				results = append(results, puAcs[rand.Intn(len(puAcs))].GetUnitInfo())
			} else {
				cfA = true
				var c []Unit
				for _, v := range Acs {
					c = append(c, v)
				}
				for _, v := range Aws {
					c = append(c, v)
				}
				results = append(results, c[rand.Intn(len(c))].GetUnitInfo())
			}
		} else {
			coS++
			coA++
			results = append(results, Bws[rand.Intn(len(Bws))].GetUnitInfo())
		}
	}

	return results
}

func drawRank(offset4 uint, offset5 uint) string {
	rate5 := math.Max(float64(offset5) + 1 - 73.0, 0) * 6.0 + 0.6
	rate4 := math.Max(float64(offset4) + 1 - 8.0, 0) * 51 + 5.1

	random := float64(rand.Intn(1000)) / 10
	if random < rate5 {
		return "S"
	} else if random < rate5 + rate4 {
		return "A"
	}
	return "B"
}

func fetchCandidate(rank string) []Unit {
	var cs []Character
	var ws []Weapon

	db := database.GetDB()
	db.Where(&Character{Rank: rank}).Find(&cs)
	db.Where(&Weapon{Rank: rank}).Find(&ws)

	var candidates []Unit
	for _, c := range(cs) {
		candidates = append(candidates, c)
	}
	for _, w := range(ws) {
		candidates = append(candidates, w)
	}
	return candidates
}

func drawPU() bool {
	c := []bool{true, false}
	return c[rand.Intn(len(c))]
}
