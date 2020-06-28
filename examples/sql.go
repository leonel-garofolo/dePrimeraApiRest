package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := "root"
	pass := "root"
	database := "de_primera_app"

	db, err := sql.Open("mysql", user+":"+pass+"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("Failed to connect database")
	}
	defer db.Close()

	dbSelect(db)
	dbSelectOne(db)

}

func dbSelect(db *sql.DB) {
	//Get all the tables from Database
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalln("Failed to query")
	}
	for rows.Next() {
		var tables string
		rows.Scan(&tables)
		log.Println(tables)
	}
}

func dbSelectOne(db *sql.DB) {
	//Get all the tables from Database
	id := 6
	row := db.QueryRow("select nombre from ligas where id_liga = ?", id)
	var nombre string
	err := row.Scan(&nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	log.Println(nombre)
}