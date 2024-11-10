package models

import (
	"fmt"
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

func (g CharacterGacha)Draw(num uint, offsetA uint, offsetS uint, fixedA bool, fixedS bool) []UnitInfo {
	// setup
	var puA_characters []Character
	var S_characters []Character
	var A_characters []Character
	var A_weapons	[]Weapon
	var B_weapons []Weapon

	db := database.GetDB()
	db.Preload("PUS").Find(&CharacterGacha{})
	db.Where([]uint{g.PUA_0ID, g.PUA_1ID}).Find(&puA_characters)
	db.Where(&Character{Rank: "S", IsLimited: false}).Not(g.PUS_ID).Find(&S_characters)
	db.Where(&Character{Rank: "A", IsLimited: false}).Not([]uint{g.PUA_0ID, g.PUA_1ID}).Find(&A_characters)
	db.Where(&Weapon{Rank: "A", IsLimited: false}).Find(&A_weapons)
	db.Where(&Weapon{Rank: "B", IsLimited: false}).Find(&B_weapons)

	fmt.Println(g.PUS_ID)
	// fmt.Println(puA_characters)
	// fmt.Println(S_characters)
	// fmt.Println(A_characters)

	// gacha
	currentOffsetS := offsetS
	currentOffsetA := offsetA
	currentFixedS := fixedS
	currentFixedA := fixedA
	var results []UnitInfo
	for i := 0; i < int(num); i++ {
		rank := drawRank(currentOffsetA, currentOffsetS)
		if rank == "S" {
			currentOffsetS = 0
			currentOffsetA = 0
			if currentFixedS || drawPU() {
				currentFixedS = false
				results = append(results, g.PUS.GetUnitInfo())
			} else {
				currentFixedS = true
				results = append(results, S_characters[rand.Intn(len(S_characters))].GetUnitInfo())
			}
		} else if rank == "A" {
			currentOffsetS++
			currentOffsetA = 0
			if currentFixedA || drawPU() {
				currentFixedA = false
				results = append(results, puA_characters[rand.Intn(len(puA_characters))].GetUnitInfo())
			} else {
				currentFixedA = true
				var c []Unit
				for _, v := range A_characters {
					c = append(c, v)
				}
				for _, v := range A_weapons {
					c = append(c, v)
				}
				results = append(results, c[rand.Intn(len(c))].GetUnitInfo())
			}
		} else {
			currentOffsetS++
			currentOffsetA++
			results = append(results, B_weapons[rand.Intn(len(B_weapons))].GetUnitInfo())
		}
	}

	return results
}

func drawRank(offset4 uint, offset5 uint) string {
	rate5 := math.Max(float64(offset5) + 1 - 73.0, 0) * 6.0 + 0.6
	rate4 := math.Max(float64(offset4) + 1 - 8.0, 0) * 51 + 5.1

	random := float64(rand.Intn(1000)) / 10
	fmt.Println(random)
	if random < rate5 {
		return "S"
	} else if random < rate5 + rate4 {
		return "A"
	}
	return "B"
}

func fetchCandidate(rank string) []Unit {
	var characters []Character
	var weapons 	 []Weapon

	db := database.GetDB()
	db.Where(&Character{Rank: rank}).Find(&characters)
	db.Where(&Weapon{Rank: rank}).Find(&weapons)

	var candidates []Unit
	for _, c := range(characters) {
		candidates = append(candidates, c)
	}
	for _, w := range(weapons) {
		candidates = append(candidates, w)
	}
	return candidates
}

func drawPU() bool {
	c := []bool{true, false}
	return c[rand.Intn(len(c))]
}
