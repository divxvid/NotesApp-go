package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	WriteJsonResponse(w, http.StatusOK, GenericResponse{
		Message: "Hello from the Golang server!",
	})
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST /signup")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST /login")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET /logout")
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetAllNotesHandler(w, r)
	} else if r.Method == "POST" {
		CreateNewNoteHandler(w, r)
	} else {
		fmt.Fprint(w, "Method not supported!")
	}
}

func NotesWithIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if r.Method == "GET" {
		GetNoteByIdHandler(id, w, r)
	} else if r.Method == "DELETE" {
		DeleteNoteByIdHandler(id, w, r)
	} else {
		fmt.Fprint(w, "Method not supported")
	}
}

func GetAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET /notes")
}

func CreateNewNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST /notes")
}

func GetNoteByIdHandler(id string, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET /notes/%s", id)
}

func DeleteNoteByIdHandler(id string, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE /notes/%s", id)
}
