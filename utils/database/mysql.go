package database

import (
	"fmt"
	"time"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/config"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitialDB(cnf config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cnf.Database.DbUser, cnf.Database.DbPass, cnf.Database.DbHost, cnf.Database.DbPort, cnf.Database.DbName)
	// dsn := fmt.Sprintf("root:root@tcp(localhost:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local")

	var db *gorm.DB
	var err error
	maxRetry := 5
	retry := 0

	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed to open database:", err.Error())
			retry++
			if retry >= maxRetry {
				log.Fatal("Max retry reached. Cannot connect to database.")
			}
			fmt.Println("Retrying connection in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	fmt.Println("database connected")
	return db
}

func TableMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.UserModels{},
		entity.OTPModels{},
		entity.CampaignModels{},
		entity.CampaignImageModels{},
		entity.TransactionModels{},
		entity.AssistantModel{},
	)
	if err != nil {
		log.Fatal("Migration table is failed", err.Error())
	} else {
		log.Info("Migration table is success")
	}
}
