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

	for rows.Next() {
		var ID, Name, BirthDate, Address, Email, Phone string
		s := &Student{ID, Name, BirthDate, Address, Email, Phone}
		if err := rows.Scan(&s.ID, &s.Name, &s.BirthDate, &s.Address, &s.Email, &s.Phone); err != nil {
			log.Fatal(err)
		}
		allStudents = append(allStudents, s)
	}
	studentsJSON, _ := json.Marshal(allStudents)
	w.Header().Set("Content-Type", "application/json")
	w.Write(studentsJSON)
}

func GetDetails(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(&Student{})
	fmt.Printf("%+v", json.NewEncoder(w).Encode(&Student{}))
}

func DeleteStudent(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, item := range students {
		if item.ID == params["id"] {
			students = append(students[:i], students[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

func EditStudent(w http.ResponseWriter, req *http.Request) {

}

func CreateStudent(w http.ResponseWriter, req *http.Request) {
	var student Student
	_ = json.NewDecoder(req.Body).Decode(&student)
	// if err != nil {
	// 	log.Fatal("Error occured while decoding request body: ", err)
	// }
	students = append(students, student)
	json.NewEncoder(w).Encode(students)
}

var students []Student

type Student struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
