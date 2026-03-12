package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jonathon-chew/URL_Shortner/app"
)

func main() {
	log.Print("Runnning...")

	if len(os.Args) >= 2 {
		app.SetCommandFlags(app.Cli(os.Args[1:]))
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./template/favicon.ico")
	})

	// API FEATURES
	// mux.HandleFunc("/api/shorten/", app.StartPage)
	mux.HandleFunc("/api/auth/", app.AuthPage)
	mux.HandleFunc("POST /api/add/", app.AddURL)

	// Simple Pages
	mux.HandleFunc("/shorten/", app.StartPage)
	// mux.HandleFunc("/", app.LogInPage)re

	// Core Feature/s
	mux.HandleFunc("/r/", app.RedirectPage)

	// Auth metods
	mux.HandleFunc("/auth/login/", app.LogInPage)
	mux.HandleFunc("/auth/signup/", app.SignUpPage)

	mux.HandleFunc("/", app.LogInPage)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
