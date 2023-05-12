package main
import (
	"encoding/json"
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
	var requestJson UserSignupRequest
	json.NewDecoder(r.Body).Decode(&requestJson)

	response := UserSignupResponse{
		Message:  "User created successfully!",
		Username: requestJson.Username,
	}
	WriteJsonResponse(w, http.StatusCreated, response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST /login")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET /logout")
}

func GetAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET /notes")
}

func CreateNewNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST /notes")
}

func GetNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "GET /notes/%s", id)
}

func DeleteNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "DELETE /notes/%s", id)
}
