package main

import (
	"fmt"
	"log"
	"net/http"
	"encode/json"
)

func ServeStudentsTable(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/views/studentstable.html")
}

func ReceiveStudentData(w http.ResponseWriter, r *http.Request) {
	r.()
	fmt.Println(r.Form)
	fmt.Println(r.FormValue("name"))
}

func main() {
	http.HandleFunc("/students_table/", ServeStudentsTable)
	http.HandleFunc("/save_student", ReceiveStudentData)
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type Student struct {
	Name      string `json: "name"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birthDate"`
	Address   string `json: "address"`
	Email     string `json:"email"`
	Phone     string `json:"string"`
}
