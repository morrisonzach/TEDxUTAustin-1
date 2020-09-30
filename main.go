package main

import (
	"html/template"
	"log"
	"net/http"
	"net/smtp"

	"github.com/bmizerany/pat"
)

//Vote is ...
type Vote struct {
	Name   string
	Email  string
	Answer string
}

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(homeHandler))
	mux.Post("/", http.HandlerFunc(voteHandler))

	log.Println("listening...")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "www/index.html", nil)
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("invalid request")
	}
	vote := Vote{
		Name:   r.FormValue("name"),
		Email:  r.FormValue("email"),
		Answer: r.FormValue("voting"),
	}
	sendVote(vote)
	render(w, "www/complete.html", nil)
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func sendVote(v Vote) {
	from := "ziemboy9@gmail.com"
	password := "Z@chary16"

	to := "hypemephotography@gmail.com"

	smtpServer := smtpServer{
		host: "smtp.gmail.com",
		port: "587",
	}

	message := []byte("From: " + from + "\n" + "To: " + to + "\n" + "Subject: Vote from " + v.Name + "!\n\nVote from " + v.Name + " with email of " + v.Email + " voted for " + v.Answer + ".")

	if err := smtp.SendMail(smtpServer.Address(), smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(message)); err != nil {
		log.Printf("smtp error: %s\n", err)
		return
	}
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
