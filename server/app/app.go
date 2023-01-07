package app

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var PGConnectionUrl = os.Getenv("POSTGRES_CONNECTION_URL")
	fmt.Println(PGConnectionUrl)
	db, err := gorm.Open(postgres.Open(PGConnectionUrl))
	if err != nil {
		panic(err)
	}

	fmt.Println(db)
	DB = db
}
