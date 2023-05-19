package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const address string = "127.0.0.1:8000"

// This is the main function
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/signup", SignupHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")
	r.HandleFunc("/notes", GetAllNotesHandler).Methods("GET")
	r.HandleFunc("/notes", CreateNewNoteHandler).Methods("POST")
	r.HandleFunc("/notes/{id}", GetNoteByIdHandler).Methods("GET")
	r.HandleFunc("/notes/{id}", DeleteNoteByIdHandler).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Server starting at ", address)
	log.Fatal(srv.ListenAndServe())
}
