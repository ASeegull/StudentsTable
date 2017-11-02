package main

import (
	"log"
	"net/http"
	//"encode/json"
)

func ServeStudentsTable(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/views/studentstable.html")
}

func main() {
	http.HandleFunc("/students_table/", ServeStudentsTable)
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// type Student struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Sex   string `json:"sex"`
// 	Email string `json:"email"`
// 	Phone string `json:"string"`
// }

// func HandleError() {
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
