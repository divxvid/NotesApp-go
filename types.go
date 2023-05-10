package main

type GenericResponse struct {
	Message string `json:"message"`
}

type UserSignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSignupResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
