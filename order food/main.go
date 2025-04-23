package main

import "fmt"
import "net/http"
import "goproject/handlers"

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/verify", handlers.VerifyHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	err := http.ListenAndServe(":8000", nil) // Start the WEB Server AT PORT 8000
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
