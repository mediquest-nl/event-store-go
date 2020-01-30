package main

import (
	"eventStore"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var msgId int = 0

const PORT = 8080

func WriteRequest(w http.ResponseWriter, r *http.Request) {
	var msg string = fmt.Sprintln("[", msgId, " ", r, "]")
	msgId++
	// write an event
	eventStore.WriteEvent(msg)
	// prepare a simple response
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte("<html><head>TEST</head><body>Hello World</body></html>"))
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		WriteRequest(w, r)
		break
		/*	case http.MethodGet:
				List(w, r)
				break
			case http.MethodDelete:
				Remove(w, r)
				break */
	default:
		log.Panicln("Method not allowed: " + r.Method)
		break
	}
}

func main() {
	log.Println("Attempting to start HTTP Server.")

	http.HandleFunc("/", HandleRequest)

	var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}
}
