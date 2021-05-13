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
	db, err := gorm.Open("postgres", "host=tacir.com port=5432 user=proxypool dbname=proxypool sslmode=disable password=proxypool")
	if err == nil {
		fmt.Println("DB connect success: ", DB.Name())
	}
	return
}
