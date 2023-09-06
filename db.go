package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("fail to connect DB : %v\n", err)
		os.Exit(1)
	}

	return db
}
