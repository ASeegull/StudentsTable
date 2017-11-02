package main

import (
	"net/http"
	//"encode/json"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", nil)
	http.HandleFunc("/students_table", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/views/studentstable.html")
	})
	// HandleError()
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
