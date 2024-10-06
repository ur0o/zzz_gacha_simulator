package models

import (
	"zzz_gacha/database"
)

func Migrate() {
	db := database.GetDB()

	// ここにマイグレーションしたいモデルを記述する
	db.AutoMigrate(&Character{})
	db.AutoMigrate(&Weapon{})
	db.AutoMigrate(&CharacterGacha{})
}