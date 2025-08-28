package main

import (
	"log"
	"net/http"
	"time"
)

var html = `<!DOCTYPE html>
<html>
  <h1>hi this is my first page</h1>
  <h2>this is the backend dev</h2>
</html>`

var jsonInfo = `{
	"name": "Tanmaya Panigrahi",
	"age": 22
}`

func main() {
	address := ":8080"
	mux := http.NewServeMux()

	// Home route (HTML)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(html))
	})

	// Users route (JSON)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonInfo))
	})

	// Server configuration
	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Start Server at", address)
	log.Fatal(s.ListenAndServe())
}
