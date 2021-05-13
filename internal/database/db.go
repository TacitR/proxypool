package database

import (
	"fmt"
	"os"
	"github.com/TacitR/proxypool/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connect() (err error) {
	dsn := "host=tacitr.com user=proxypool password=proxypool dbname=proxypool port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println("DB connect success: ", DB.Name())
	}
	return
}
