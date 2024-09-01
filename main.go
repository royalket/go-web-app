package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/index.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/ci-cd-best-practices.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/cloud-infrastructure-optimization.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/python-data-engineering.html")
}

func main() {

	http.HandleFunc("/index", homePage)
	http.HandleFunc("/ci-cd-best-practices", coursePage)
	http.HandleFunc("/cloud-infrastructure-optimization", aboutPage)
	http.HandleFunc("/python-data-engineering", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
