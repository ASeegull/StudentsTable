package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type Student struct {
	ID           string `json:"id"`
	Student_name string `json:"student_name"`
	Birth_date   string `json:"birthdate"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}

func main() {
	dbconn, err := sql.Open("mysql", "root:korrasami@tcp(localhost:3306)/students")
	if err != nil {
		log.Println(err)
	}
	db = dbconn
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", GetMainPage).Methods("GET")
	router.HandleFunc("/students", GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", GetDetails).Methods("GET")
	router.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")
	router.HandleFunc("/students/{id}", EditStudent).Methods("PUT")
	router.HandleFunc("/new_student", CreateStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetMainPage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `Welcome!
		GET    "/students/{id}" to get details about specifeic student
		POST   "/new_student" to add new student
		DELETE "/students/{id}" to delete specific student
		PUT    "/students/{id}" to update student's info
	`)
}

func GetAllStudents(w http.ResponseWriter, req *http.Request) {
	var allStudents []*Student

	rows, err := db.Query("SELECT * from students;")
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var ID, Student_name, Birth_date, Address, Email, Phone string
		s := &Student{ID, Student_name, Birth_date, Address, Email, Phone}
		if err := rows.Scan(&s.ID, &s.Student_name, &s.Birth_date, &s.Address, &s.Email, &s.Phone); err != nil {
			log.Fatal(err)
		}
		allStudents = append(allStudents, s)
	}
	studentsJSON, _ := json.Marshal(allStudents)
	w.Header().Set("Content-Type", "application/json")
	w.Write(studentsJSON)
}

func GetDetails(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var ID, Student_name, Birth_date, Address, Email, Phone string
	s := &Student{ID, Student_name, Birth_date, Address, Email, Phone}
	err := db.QueryRow("SELECT * FROM students WHERE id=?;", id).Scan(&s.ID, &s.Student_name, &s.Birth_date, &s.Address, &s.Email, &s.Phone)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	sJSON, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	w.Write(sJSON)
}

func DeleteStudent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	stmt, _ := db.Prepare("DELETE from students WHERE id=?")
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal("Error occured while deleting from db: ", err)
	}
	fmt.Println(res)
}

func EditStudent(w http.ResponseWriter, req *http.Request) {
	var updateValues map[string]string
	vars := mux.Vars(req)
	id := vars["id"]
	err := json.NewDecoder(req.Body).Decode(&updateValues)
	if err != nil {
		log.Fatal("Error occured while decoding request body: ", err)
	}
	var query string
	for key, value := range updateValues {
		query += fmt.Sprintf(" %v = '%v',", key, value)
	}
	query = "UPDATE students SET " + query[:len(query)-1] + " WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("Error occured while creating insert statement: ", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal("Error occured while inserting into db: ", err)
	}
	fmt.Printf("%v", res)
}

func CreateStudent(w http.ResponseWriter, req *http.Request) {
	var ns Student
	err := json.NewDecoder(req.Body).Decode(&ns)
	if err != nil {
		log.Fatal("Error occured while decoding request body: ", err)
	}
	stmt, err := db.Prepare("INSERT INTO students VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error occured while creating insert statement: ", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(ns.ID, ns.Student_name, ns.Birth_date, ns.Address, ns.Email, ns.Phone)
	if err != nil {
		log.Fatal("Error occured while inserting into db: ", err)
	}
	fmt.Printf("%v", res)
}
