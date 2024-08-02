package main

type ApiError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
	Data  interface{} `json:"data"`
	Error ApiError    `json:"error"`
}