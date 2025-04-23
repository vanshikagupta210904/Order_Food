package handlers

import (
	"goproject/config"
	"html/template"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	phone := r.URL.Query().Get("phone")
	if phone == "" {
		http.Error(w, "Phone number is required", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		// Get the user's name and email from the form
		name := r.FormValue("name")
		email := r.FormValue("email")

		if name == "" || email == "" {
			http.Error(w, "Name and Email are required", http.StatusBadRequest)
			return
		}

		// Connect to DB
		db, err := config.InitDB()
		if err != nil {
			http.Error(w, "DB connection error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Insert new user into the database
		_, err = db.Exec("UPDATE users SET name = ?, email = ?, registration_date = NOW() WHERE phone_no = ?", name, email, phone)
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// Redirect to login or dashboard after signup
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// Render the signup page
	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Pass phone number to the template
	tmpl.Execute(w, struct {
		Phone string
	}{Phone: phone})
}
