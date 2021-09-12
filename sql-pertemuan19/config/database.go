package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormDatabaseConn() (*gorm.DB, error) {
	dsn := "root:rakamin@tcp(127.0.0.1:3306)/shipment?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Cannot connect db", err)
		return nil, err
	}
	fmt.Println("Connect to database.")

	return db, nil
}
