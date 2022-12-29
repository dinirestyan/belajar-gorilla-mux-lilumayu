//file main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Selamat datang di skillplus</h1>")
	} else if r.URL.Path == "/aboutyou" {
		fmt.Fprint(w, "<h1>Skillplus adalah web tutorial seputar informasi teknologi</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Halaman yang dicari tidak ditemukan</h1>")
	}
}

func main() {
	r := mux.NewRouter()
	d := mux.NewRouter()
	d.HandleFunc("/", handlerFunc)
	d.HandleFunc("/aboutyou", handlerFunc)
	d.HandleFunc("/aboutme", handlerFunc)
	r.HandleFunc("/", handlerFunc)
	r.HandleFunc("/aboutus", handlerFunc)
	r.PathPrefix("/dini").Handler(http.StripPrefix("/dini", d))
	log.Print("running")
	log.Fatal(http.ListenAndServe(":3000", r))
}
