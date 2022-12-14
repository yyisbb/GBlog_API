package config

import (
	"gblog_api/global"
	"gblog_api/models"
	"log"
)

func DBAutoCreate() {
	err := global.GlobalMysql.AutoMigrate(&models.Article{},
		&models.Setting{},
		&models.Category{},
		&models.User{})
	if err != nil {
		log.Println("[DBAutoCreate]")
	}
}
