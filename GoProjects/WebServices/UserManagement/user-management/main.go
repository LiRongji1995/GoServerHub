package main

import (
	"log"
	"net/http"
	"user-management/handlers"
	"user-management/middleware"
)

func main() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetUsers)))
	http.Handle("/users/", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateUser)))
	http.Handle("/roles/assign", middleware.AuthMiddleware(http.HandlerFunc(handlers.AssignRole)))
	http.Handle("/report", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetReport)))
	http.Handle("/security/https", middleware.AuthMiddleware(http.HandlerFunc(handlers.EnableHTTPS)))

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
