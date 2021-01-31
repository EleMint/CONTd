package controllers

import (
	"log"
	"net/http"
)

// Init initalizes each service controller
func Init() {

	log.Println("Initializing base controller")
	InitBaseController()
	http.HandleFunc("/", BaseHTTPHandleFunc)

	log.Println("Initializing user controller")
	InitUserController()
	http.HandleFunc("/user", UserHTTPHandleFunc)
	http.HandleFunc("/users", UserHTTPHandleFunc)
}
