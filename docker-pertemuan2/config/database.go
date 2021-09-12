package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() (*gorm.DB, error) {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/mysql"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Cannot connect db", err)
		return nil, err
	}

	fmt.Println("Connect to database.")

	return db, nil
}
