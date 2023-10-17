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
	// deleteStudent("9")
	// students := getAllStudent()
	// for _, student := range students {
	// 	fmt.Println(student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	// }
	// fmt.Println(getStudentById(7))
	students := searchBy("ik", "2000-11-30")
	for _, student := range students {
		fmt.Println(student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	}
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

func getAllStudent() []entity.Student {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_student;"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := scanStudent(rows)

	return students
}

func scanStudent(rows *sql.Rows) []entity.Student {
	students := []entity.Student{}
	var err error

	for rows.Next() {
		student := entity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.BirthDate, &student.Gender)
		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return students
}

func getStudentById(id int) entity.Student {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM mst_student WHERE id = $1;"

	student := entity.Student{}
	err = db.QueryRow(sqlStatement, id).Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.BirthDate, &student.Gender)
	if err != nil {
		panic(err)
	}
	return student
}

func searchBy(name string, birthDate string) []entity.Student {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_student WHERE name LIKE $1 AND birth_date = $2;"

	rows, err := db.Query(sqlStatement, "%"+name+"%", birthDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	students := scanStudent(rows)
	return students
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
