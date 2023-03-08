package models

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open db ->", err.Error())
		os.Exit(1)
	}

	err = database.AutoMigrate(&Room{})
	if err != nil {
		return
	}

	DB = database
}
