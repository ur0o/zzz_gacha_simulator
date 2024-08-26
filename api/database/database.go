package database

import (
	"os"

	"fmt"
	"github.com/go-yaml/yaml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Protocol string `yaml:"protocol"`
	Host 		 string `yaml:"host"`
	Port 		 string `yaml:"port"`
	Name 		 string `yaml:"name"`
	User 		 string `yaml:"user"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

var db *gorm.DB

func InitDB() {
	config := readConfig()
	db = openDB(config)
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	d, _ := db.DB()
	d.Close()
}

func readConfig() (config *DBConfig) {
	b, _ := os.ReadFile("database/database.yml")
	yaml.Unmarshal(b, &config)

	return
}

func openDB(config *DBConfig) *gorm.DB {
	eu := config.User + ":" + config.Password + "@" + config.Protocol + "(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=" + config.Charset + "&parseTime=True&loc=Local"
	fmt.Println(eu)
	d, err := gorm.Open(mysql.Open(eu), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return d
}