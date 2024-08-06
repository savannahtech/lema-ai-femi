package repositories

import (
	"fmt"
	"github.com/djfemz/savannahTechTask/api/models"
	"github.com/djfemz/savannahTechTask/api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
)

func ConnectToDatabase() (*gorm.DB, error) {
	port, err := strconv.ParseUint(os.Getenv(utils.DATABASE_PORT), 10, 64)
	if err != nil {
		log.Fatal("Error reading port: ", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Africa/Lagos sslmode=disable", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	err = db.AutoMigrate(&models.Commit{}, &models.GithubRepository{})
	if err != nil {
		log.Fatal("error migrating: ", err)
	}
	err = db.AutoMigrate(&models.Author{})
	if err != nil {
		log.Fatal("error migrating: ", err)
	}
	return db, nil
}
