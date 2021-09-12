package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:myexamplepassword@tcp(mysql-docker:3307)/mysql")
	if err != nil {
		fmt.Println("Error DB:", err.Error())
		return
	}
	defer db.Close()

	fmt.Println("Connect to database.")

	name := os.Getenv("NAME")
	port := os.Getenv("PORT")

	http.HandleFunc("/", hello)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	fmt.Println("Hello World, dari port", port, name)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
