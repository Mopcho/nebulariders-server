package main

type ApiError struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type AuthService struct {
	Users map[string]*User
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type ApiResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error *ApiError   `json:"error,omitempty"`
}