package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseInit() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:rakamin@tcp(127.0.0.1:3306)/animal")

	if err != nil {
		// panic(err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(100) //Membatasi open koneksi ke database
	db.SetMaxIdleConns(80)  //Membatasi idle koneksi ke database

	fmt.Println("Connected to database")

	return db, nil
}
