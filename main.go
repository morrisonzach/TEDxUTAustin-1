package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func main() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":5500", nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	r.ParseForm()

	fmt.Fprintf(w, "hello, this is the form: "+r.Form.Get("name"))
	from := "ziemboy9@gmail.com"
	pass := "Z@chary14"

	to := []string{
		"hypemephotography@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(r.Form.Get("voting"))
	auth := smtp.PlainAuth("", from, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("email sent successfully!")

}
