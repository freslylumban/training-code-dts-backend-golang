package handler

import (
	"fmt"
	"schema-modelling-gorm/config"
	"schema-modelling-gorm/model"
)

func InsertPerson() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	person := model.Person{Name: "Fresly", Email: "freslylu@gmail.com"}
	db.Create(&person)

	fmt.Println("Success insert using ORM")
}

func InsertUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	// user := model.User{Name: "Adrian", Email: "adrian@gmail.com", Address: "Bandung"}
	user := model.User{Name: "Julian", Email: "julian@gmail.com", Address: "Surabaya"}
	db.Create(&user)

	fmt.Println("Insert user successfully.")
}

func SelectAllUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var users []model.User

	result := db.Find(&users)

	fmt.Println("&users", &users)
	fmt.Println("Result.RowsAffected", result.RowsAffected)
	fmt.Println("Result", result)
	fmt.Println("Select all users successfully.")
}

func SelectUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	result := db.First(&user, 2)

	fmt.Println("&users", &user)
	fmt.Println("Result.RowsAffected", result.RowsAffected)
	fmt.Println("Result", result)
	fmt.Println("Select all users successfully.")
}

func SelectUserLooping() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	// result := db.First(&user)
	// db.First(&user, "name = ?", "Adrian")

	result, err := db.Model(&model.User{}).Rows()

	if err != nil {
		fmt.Errorf("Cannot scan row", err.Error())
	}

	for result.Next() {
		db.ScanRows(result, &user)
		fmt.Println("name:", user.Name)
		fmt.Println("email:", user.Email)
	}

	fmt.Println("Result:", &result)
	fmt.Println("Select users with looping successfully.")
}

func UpdateUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	db.Model(&user).Where("id = ?", 1).Update("name", "Hordon")

	fmt.Println("Updated user successfully.")
}

func DeleteUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	db.Where("id = ?", 3).Delete(&user)

	fmt.Println("Deleted user successfully.")
}
