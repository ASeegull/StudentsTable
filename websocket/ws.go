package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.Path("/counter").HandlerFunc(counter)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}

func counter(w http.ResponseWriter, req *http.Request) {
	wsconn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("COuld not establish websocket, error: ", err)
	}
	fmt.Println("Connection established")
	defer wsconn.Close()

	type Counter struct {
		Count int `json:"count"`
	}
	c := new(Counter)

	for {
		messageType, msg, err := wsconn.ReadMessage()
		fmt.Println(messageType, string(msg))

		if err != nil {
			log.Println("Could not read message from websocket, error", err)
			return
		}

		continueCount := true
		if messageType == websocket.CloseMessage {
			log.Println("closing websocket... ")
			continueCount = false
			break
		}

		if string(msg) == "reset" {
			c.Count = 0
		}

		for continueCount {
			time.Sleep(1 * time.Second)
			c.Count++
			databytes, err := json.Marshal(c)
			if err = wsconn.WriteMessage(messageType, databytes); err != nil {
				log.Println("Could not write message to websocket, error", err)
				return
			}
		}
	}
}
