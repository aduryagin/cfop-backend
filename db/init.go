package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Instance *gorm.DB
var err error

func Init() {
	Instance, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=cfop sslmode=disable")

	if err != nil {
		panic(err)
	}

	Instance.AutoMigrate(&Group{}, &Subgroup{}, &Algorithm{})
}
