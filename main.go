package main

import (
	"database/sql"
	"fmt"

	"database-example/entity"

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

	// student := entity.Student{Id: 9, Name: "Siti kocak", Email: "sitikocak@yahoo.com", Address: "Jakarta Selatan", BirthDate: time.Date(2000, 11, 20, 0, 0, 0, 0, time.Local), Gender: "F"}

	// addStudent(student)
	// updateStudent(student)
	deleteStudent("9")
}

func addStudent(student entity.Student) {
	db := connectDb()
	defer db.Close()
	var err error

	// Untuk menghindari SQL Injection kita dapat menggunakan $

	sqlStatement := "INSERT INTO mst_student (id, name, email, address, birth_date, gender) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Successfully Insert Data!")
	}

}

func updateStudent(student entity.Student) {
	db := connectDb()
	defer db.Close()
	var err error

	// Untuk menghindari SQL Injection kita dapat menggunakan $

	sqlStatement := "UPDATE mst_student SET name = $2, email = $3, address = $4, birth_date = $5, gender = $6 WHERE id = $1;"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Successfully Update Data!")
	}

}

func deleteStudent(id string) {
	db := connectDb()
	defer db.Close()
	var err error

	// Untuk menghindari SQL Injection kita dapat menggunakan $

	sqlStatement := "DELETE FROM mst_student WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Successfully Delete Data!")
	}
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// Close the database connection when we'r done with it.
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Connected")
	}
	return db
}
