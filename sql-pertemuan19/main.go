package main

import (
	"fmt"
	"migrate-gorm/config"
	"migrate-gorm/model"
)

func main() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error:", err.Error())
	}

	db.AutoMigrate(&model.Customer{}, &model.Order{}, &model.Shipment{})

}
