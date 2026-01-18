package main

import (
	"auth/internal/auth"
	database "auth/internal/database"
	http_auth "auth/internal/http"
	"log"
	"net/http"
)

func main() {
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := database.NewUserRepository(db)
	tokenRepo := database.NewRefreshTokenRepo(db)
	jwt := auth.NewJWTService("super-secret")

	auth := http_auth.NewAuthService(userRepo, tokenRepo, jwt)
	h := http_auth.NewHandler(auth)

	http.HandleFunc("/register", h.Register)
	http.HandleFunc("/login", h.Login)

	log.Println("Auth server on: 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}