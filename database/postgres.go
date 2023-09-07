package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Preman struct {
	gorm.Model
	// Id  int `json:"id"`
	Nama string `json:"nama"`
}

func (Preman) TableName() string {
    return "preman"
}

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "db_preman"
	dbUser := "postgres"
	password := "0257"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Preman{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
