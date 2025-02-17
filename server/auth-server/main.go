package main

import (
	"log"
	"net/http"

	"github.com/Mopcho/nebulariders-server/common/mopHttp"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func (s *AuthService) handleVerifyToken(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	parsedToken, err := verifyToken(token)
	if err != nil {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message: "Invalid Token"}}, http.StatusUnauthorized)
		return
	}

	subject, err := parsedToken.Claims.GetSubject()
	if err != nil {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"Failed to get token subject"}}, http.StatusInternalServerError)
		return
	}

	// Check if user exists with this sub
	userExists := false
	for id := range s.Users {
		user := s.Users[id]
		if user.ID == subject {
			userExists = true
		}
	}

	if !userExists {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"User for this token does not exist anymore"} }, http.StatusUnauthorized)
		return
	}

	mopHttp.SendResponse(w, mopHttp.ApiResponse{Data: "ok" }, http.StatusOK)
}

func (s *AuthService) handleLogin(w http.ResponseWriter, r *http.Request) {
	var decodedBody = LoginRequestBody{}

	if err := mopHttp.GetJsonBody(r.Body, &decodedBody); err != nil {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"Could not parse request body"}}, http.StatusInternalServerError)
		return
	}

	var user, ok = s.Users[decodedBody.Email]
	if !ok {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"Invalid Credentials"}}, http.StatusUnauthorized)
		return
	}
	
	if 	ok := CheckPasswordHash(decodedBody.Password, user.Password); !ok {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"Invalid Credentials"}}, http.StatusUnauthorized)
		return
	}

	token, err := createToken(user.ID, user.Username)
	if err != nil {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"Internal Server Error"} }, http.StatusInternalServerError)
		return
	}

	type LoginResponseData struct { Token string `json:"token"`}
	mopHttp.SendResponse(w, mopHttp.ApiResponse{Data: LoginResponseData{Token: token} }, http.StatusOK)
}

func (s *AuthService) handleRegister(w http.ResponseWriter, r *http.Request) {
	var decodedBody = RegisterRequestBody{}
	
	if err := mopHttp.GetJsonBody(r.Body, &decodedBody); err != nil {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"Could not parse request body"}}, http.StatusInternalServerError)
	}
	
	if err := validateUserRegisterData(decodedBody); err != nil {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message: err.Error()}}, http.StatusBadRequest)
		return
	}
	
	if _, ok := s.Users[decodedBody.Email]; ok {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"User with this email already exists"}}, http.StatusConflict)
		return
	}

	hashedPass, err := hashPassword(decodedBody.Password)
	if err != nil {
		mopHttp.SendResponse(w, mopHttp.ApiResponse{Error: &mopHttp.ApiError{Message:"Failed hashing password"}}, http.StatusInternalServerError)
		return
	}

	s.Users[decodedBody.Email] = &User{ID: uuid.New().String(),Username: decodedBody.Username, Email: decodedBody.Email, Password: hashedPass}
	mopHttp.SendResponse(w, mopHttp.ApiResponse{Data: "ok"}, http.StatusOK)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var authService = AuthService{Users: make(map[string]*User)} 
	mopHttp.Post("/api/auth/login", authService.handleLogin)
	mopHttp.Post("/api/auth/register", authService.handleRegister)
	mopHttp.Post("/api/auth/verifyToken", authService.handleVerifyToken)
	log.Println("Starting auth server on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

