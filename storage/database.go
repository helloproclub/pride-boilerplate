package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func initDB(
	dbHost string,
	dbPort string,
	dbUser string,
	dbName string,
	dbPass string,
) *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(err)
	}

	return db
}
