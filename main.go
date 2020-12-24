package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int32
	Money float64
	Hobbies []string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	person := User{
		Name: "Bob",
		Age: 25,
		Money: 1234.5,
		Hobbies: []string{"Football", "Running", "Basketball"},
	}

	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, person)
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
