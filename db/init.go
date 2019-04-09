package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Instance - db instance
var Instance *gorm.DB
var err error

// Init - db init
func Init() {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=database port=5432 user=postgres dbname=cfop sslmode=disable"
	}

	Instance, err = gorm.Open("postgres", databaseURL)

	if err != nil {
		panic(err)
	}

	Instance.AutoMigrate(&Group{}, &Subgroup{}, &Algorithm{})
}
