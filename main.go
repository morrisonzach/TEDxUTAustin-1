package main

import (
	"fmt"
	"log"
	"net/http"
)

//Vote is ...
type Vote struct {
	Name   string
	Email  string
	Answer string
}

func vote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("invalid request")
	}
	vote := Vote{
		Name:   r.FormValue("name"),
		Email:  r.FormValue("email"),
		Answer: r.FormValue("voting"),
	}
	fmt.Fprintf(w, "<h1>thanks for the vote "+vote.Name+"</h1><p>"+vote.Email+" "+vote.Answer+"</p>")
}

func main() {
	http.HandleFunc("/vote", vote)
	// http.HandleFunc("/results", results)
	fmt.Println("printing to show server is on")
	log.Fatal(http.ListenAndServe("127.0.0.1:5500", nil))
}
