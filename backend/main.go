package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "home.html")
}

func serveRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	room := vars["room"]
	if room == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, room)
}

func main() {
	// Check to see if PORT has been set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	flag.Parse()
	go h.run()
	r := mux.NewRouter()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})

	r.HandleFunc("/{room}/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
	http.Handle("/", r)
	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on PORT:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
