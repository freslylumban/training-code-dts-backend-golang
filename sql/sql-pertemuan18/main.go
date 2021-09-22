package main

import (
	"fmt"
	"schema-modelling-gorm/config"
	"schema-modelling-gorm/handler"
	"schema-modelling-gorm/model"
)

func main() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error", err.Error())
	}

	db.AutoMigrate(&model.Person{}, &model.User{})
	// handler.InsertPerson()
	// handler.InsertUser()
	handler.SelectAllUser()
	// handler.SelectUserLooping()
	// handler.SelectUser()
	// handler.UpdateUser()
	// handler.DeleteUser()
}
