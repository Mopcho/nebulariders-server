package mopHttp

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Wrapper(handler func(http.ResponseWriter, *http.Request), method string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			SendResponse(w, ApiResponse{Error: &ApiError{Message: "Method not allowed"}}, http.StatusNotFound)
			return
		}

		handler(w, r)
	}
}

func Post(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, Wrapper(handler, "POST"))
}

func Get(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, Wrapper(handler, "GET"))
}

func Put(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, Wrapper(handler, "PUT"))
}

func Delete(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, Wrapper(handler, "DELETE"))
}

func Patch(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, Wrapper(handler, "PATCH"))
}

func SendResponse(w http.ResponseWriter, data ApiResponse, status int) {
	var bytes, err = json.Marshal(data)

	if err != nil {
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(status)
	w.Write(bytes)
}

func GetJsonBody(body io.ReadCloser, v any) error {
	var err = json.NewDecoder(body).Decode(&v)
	if err != nil {
		return errors.New("failed decoding request body")
	}

	return nil
}

type ApiResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error *ApiError   `json:"error,omitempty"`
}

type ApiError struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}