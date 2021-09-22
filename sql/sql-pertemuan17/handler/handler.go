package handler

import (
	"connect-and-interacting-database/config"
	"fmt"
)

func InsertAnimal() {
	db, err := config.DatabaseInit()

	if err != nil {
		// fmt.Errorf("Error connect db ", err.Error())
		fmt.Println("Error:", err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into animal(name) values (?)", "Panda")
	if err != nil {
		// fmt.Errorf("Error insert db ", err.Error())
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Inserted successfully!")
}

func UpdateAnimal() {
	db, err := config.DatabaseInit()

	if err != nil {
		// fmt.Errorf("Error connect db ", err.Error())
		fmt.Println("Error:", err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("update animal set name=? where id=?", "Sapi", 4)
	if err != nil {
		// fmt.Errorf("Error insert db ", err.Error())
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Updated successfully!")
}

func DeleteAnimal() {
	db, err := config.DatabaseInit()

	if err != nil {
		// fmt.Errorf("Error connect db ", err.Error())
		fmt.Println("Error:", err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("delete from animal where id=?", 5)
	if err != nil {
		// fmt.Errorf("Error insert db ", err.Error())
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Deleted successfully!")
}
