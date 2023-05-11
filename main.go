package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const address string = "127.0.0.1:8000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbSvc, err := NewMongoDBService()
	if err != nil {
		log.Fatal(err)
	}
	defer dbSvc.Close()

	names, err := dbSvc.GetCollectionNames()
	if err != nil {
		log.Fatal("Error in GetCollection Names: ", err)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	ups, err := dbSvc.GetAllUserPasses()
	if err != nil {
		log.Fatal("Error in GetAllUserPasses Names: ", err)
	}

	for _, up := range ups {
		fmt.Println(up)
	}

	up := UserPasses{
		Username: "GolangTest2",
		Password: "hehehe",
	}

	id, err := dbSvc.CreateUser(&up)
	if err != nil {
		log.Fatal("Error in Create User: ", err)
	}
	fmt.Println("Create succeeded with id: ", id)

	ups, err = dbSvc.GetAllUserPasses()
	if err != nil {
		log.Fatal("Error in GetAllUserPasses Names: ", err)
	}

	for _, up := range ups {
		fmt.Println(up)
	}
}

func tempFunc() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/signup", SignupHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")
	r.HandleFunc("/notes", NotesHandler).Methods("GET", "POST")
	r.HandleFunc("/notes/{id}", NotesWithIdHandler).Methods("GET", "DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Server starting at ", address)
	log.Fatal(srv.ListenAndServe())
}
