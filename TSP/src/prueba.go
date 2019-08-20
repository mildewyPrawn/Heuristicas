package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "../base/tsp.db")
	rows, _ := database.Query("SELECT id, name FROM cities")
	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(strconv.Itoa(id) + ": " + name)
	}
}
