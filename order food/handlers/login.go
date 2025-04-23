package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"goproject/otp"
)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		phone := r.FormValue("phone")
		if phone == "" {
			http.Error(w, "Phone number is required", http.StatusBadRequest)
			return
		}

		otpCode := fmt.Sprintf("%06d", rand.Intn(1000000))
		otp.SetOTP(phone, otpCode)
		
		fmt.Printf("Sending OTP %s to phone %s\n", otpCode, phone)
		http.Redirect(w, r, "/verify?phone="+phone, http.StatusSeeOther)
		return
	}

	tmpl, _ := template.ParseFiles("templates/login.html")
	tmpl.Execute(w, nil)
}
