package db

import (
	"fmt"
	"os"

	"github.com/bsalpsingh/frontRow/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	HOST := os.Getenv("HOST")
	USER := os.Getenv("USER")
	PW := os.Getenv("PASSWORD")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", HOST, USER, PW, DBNAME, PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Event{}, &models.Performer{}, &models.Venue{}, &models.Ticket{}, &models.Booking{})
	return db, err
}
