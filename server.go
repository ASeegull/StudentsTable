package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Student struct {
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func ServeStudentsTable(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/views/studentstable.html")
}

func ReceiveStudentData(w http.ResponseWriter, r *http.Request) {
	student := Student{}
	defer r.Body.Close()
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error occered while reading request body: ", err)
	}
	err = json.Unmarshal(jsn, &student)
	if err != nil {
		log.Fatal("Error occured while decoding request body: ", err)
	}
	SaveStudent(&student)
}

func SaveStudent(student *Student) {
	studentJSON, _ := json.Marshal(student)
	err := ioutil.WriteFile("students.json", studentJSON, 0644)
	if err != nil {
		log.Fatal("Error occured while writing to JSON file: ", err)
	}
	fmt.Printf("%+v", studentJSON)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/students_table/", ServeStudentsTable)
	http.HandleFunc("/save_student", ReceiveStudentData)
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
