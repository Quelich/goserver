package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Register static files path
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register router
	http.Handle("/", router)

	// Register Handlers
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/about", aboutHandler)

	// Start listening
	listen()
}

func listen() {
	log.Printf("Listening on %s%s...\n", HOST, PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.WriteHeader(http.StatusOK)
	p := "." + r.URL.Path
	if p == "./" { /*root directory*/
		p = buildHtml(r.URL.Path + "index")
	}
	http.ServeFile(w, r, p)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	p := buildHtml(r.URL.Path)
	http.ServeFile(w, r, p)
}
