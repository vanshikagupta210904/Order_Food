package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"goproject/config"
	"goproject/otp"
)

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")

	if r.Method == http.MethodPost {
		enteredOTP := r.FormValue("otp")

		// Verify OTP
		if !otp.VerifyOTP(phone, enteredOTP) {
			http.Error(w, "Invalid or expired OTP", http.StatusUnauthorized)
			return
		}

		// Connect to DB
		db, err := config.InitDB()
		if err != nil {
			http.Error(w, "DB connection error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var userID int
		// Check if the user exists
		err = db.QueryRow("SELECT user_id FROM users WHERE phone_no = ?", phone).Scan(&userID)

		if err == sql.ErrNoRows {
			// If user not found, redirect to signup
			http.Redirect(w, r, "/signup?phone="+phone, http.StatusSeeOther)
			return
		} else if err != nil {
			http.Error(w, "DB query error", http.StatusInternalServerError)
			return
		} else {
			// User found, log in
			fmt.Printf("User %d logged in\n", userID)
		}

		// Redirect to dashboard (or home)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// Display verify page with phone number
	tmpl, err := template.ParseFiles("templates/verify.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, struct {
		Phone string
	}{Phone: phone})
}
