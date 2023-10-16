package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "stanners2020"
	dbname   = "enigmacamp"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	addStudent("7", "Kiko", "kiko@yahoo.com", "Bogor", "1998-03-20", "M")
}

func addStudent(id, name, email, address, birthDate, gender string) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close() // Close the database connection when we'r done with it.
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Connected")
	}
	// Create a new table in the database called 'users'.

	sqlStatement := "INSERT INTO mst_student (id, name, email, address, birth_date, gender) VALUES (" + id + "," + "'" + name + "'" + "," + "'" + email + "'" + "," + "'" + address + "'" + "," + "'" + birthDate + "'" + "," + "'" + gender + "'" + ");"

	_, err = db.Exec(sqlStatement)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Successfully Insert Data!")
	}

}
