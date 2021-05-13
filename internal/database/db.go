package database

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/TacitR/proxypool/config"
)

var DB *gorm.DB

func connect() (err error) {
	dsn := "user=proxypool password=proxypool dbname=proxypool port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	if url := config.Config.DatabaseUrl; url != "" {
		s3:=dsn+"host="+url
	}
	if url := os.Getenv("DATABASE_URL"); url != "" {
		s3:=dsn+"host="+url
	}
	DB, err = gorm.Open(postgres.Open(s3), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err == nil {
		fmt.Println("DB connect success: ", DB.Name())
	}
	return
}
