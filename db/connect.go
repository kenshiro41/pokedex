package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

const (
	host     = "pokemon_db"
	port     = 5432
	user     = "ken41"
	password = "ken41"
	dbname   = "pokemon"
)

func Connect() *gorm.DB {
	dbConnect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", dbConnect)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbConnect)
	db.SingularTable(true)
	return db
}
