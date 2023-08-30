package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
	)(http.DefaultServeMux)

	http.ListenAndServe(":8000", corsHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request:", r.Method, r.URL.Path)
	if r.Method == http.MethodPost {
		response := randWord()
		fmt.Println(response)
		fmt.Fprint(w, response)
	}
}

func words() [20]string {
	var words [20]string

	words[0] = "Yes, definitely"
	words[1] = "No, not this time"
	words[2] = "Outlook is hazy, try again"
	words[3] = "It's a possibility"
	words[4] = "Absolutely yes"
	words[5] = "Can't say for sure"
	words[6] = "Ask again later"
	words[7] = "No doubt about it"
	words[8] = "Most likely"
	words[9] = "Better not tell you now"
	words[10] = "Not looking good"
	words[11] = "Yes, but with caution"
	words[12] = "Signs point to yes"
	words[13] = "Unlikely"
	words[14] = "You can count on it"
	words[15] = "Reply hazy, try again"
	words[16] = "Definitely no"
	words[17] = "Chances are slim"
	words[18] = "Ask someone else"
	words[19] = "It's up to you"

	return words
}

func rNum() int {
	randomNumber := rand.Intn(20)
	return randomNumber
}

func randWord() string {
	var words = words()
	var rNum = rNum()
	return words[rNum]
}
