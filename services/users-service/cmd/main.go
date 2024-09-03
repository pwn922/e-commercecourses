package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer mysecrettoken" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}


func main() {
	log.Println("Server starting...")

	router := mux.NewRouter()

	router.HandleFunc("/", loginHandler)

	router.Use(authMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router))
}