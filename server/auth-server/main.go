package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func (s *AuthService) handleVerifyToken(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "GET") {
		token := r.URL.Query().Get("token")
		parsedToken, err := verifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			sendResponse(w, ApiResponse{Error: ApiError{Message: "Invalid Token"}})
			return
		}

		subject, err := parsedToken.Claims.GetSubject()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			sendResponse(w, ApiResponse{Error: ApiError{Message:"Failed to get token subject"}})
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
			w.WriteHeader(http.StatusUnauthorized)
			sendResponse(w, ApiResponse{Error: ApiError{Message:"User for this token does not exist anymore"} })
			return
		}

		sendResponse(w, ApiResponse{Data: "ok" })
		return
	}
}

func (s *AuthService) handleLogin(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		var decodedBody = LoginRequestBody{}
		err := getJsonBody(w, r.Body, &decodedBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			sendResponse(w, ApiResponse{Error: ApiError{Message:"Could not parse request body"}})
			return
		}

		var user, ok = s.Users[decodedBody.Email]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			sendResponse(w, ApiResponse{Error: ApiError{Message:"Invalid Credentials"}})
			return
		}
		
		if user.Password != decodedBody.Password {
			w.WriteHeader(http.StatusUnauthorized)
			sendResponse(w, ApiResponse{Error: ApiError{Message:"Invalid Credentials"}})
			return
		}
		token, err := createToken(user.ID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			sendResponse(w, ApiResponse{Error: ApiError{Message:"Internal Server Error"} })
			return
		}

		type LoginResponseData struct { Token string `json:"token"`}
		sendResponse(w, ApiResponse{Data: LoginResponseData{Token: token} })
		return
	}
}

func (s *AuthService) handleRegister(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		var decodedBody = RegisterRequestBody{}
		err := getJsonBody(w, r.Body, &decodedBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			sendResponse(w, ApiResponse{Error: ApiError{Message:"Could not parse request body"}})
		}

		var _, ok = s.Users[decodedBody.Email]
		if ok {
			w.WriteHeader(http.StatusConflict)
			sendResponse(w, ApiResponse{Error: ApiError{Message:"User with this email already exists"}})
			return
		}

		s.Users[decodedBody.Email] = &User{ID: uuid.New().String(),Username: decodedBody.Username, Email: decodedBody.Email, Password: decodedBody.Password}
		sendResponse(w, ApiResponse{Data: "ok"})
	}	
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var authService = AuthService{Users: make(map[string]*User)} 
	http.HandleFunc("/api/auth/login", authService.handleLogin)
	http.HandleFunc("/api/auth/register", authService.handleRegister)
	http.HandleFunc("/api/auth/verifyToken", authService.handleVerifyToken)
	log.Println("Starting auth server on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

