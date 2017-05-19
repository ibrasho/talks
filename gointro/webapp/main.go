package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/form", handleFormSubmit)
	http.HandleFunc("/page", handleTemplatePage)
	http.HandleFunc("/", handleHome)

	fmt.Println("Listening on port 8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if name == "" {
		w.Write([]byte("Hello world"))
	} else {
		fmt.Fprintf(w, "Hello, %v", name)
	}
}

func handleTemplatePage(w http.ResponseWriter, r *http.Request) {
	p := struct {
		Title string
		Body  string
	}{"Page Title", "Page Body"}

	t, _ := template.ParseFiles("edit.html")

	t.Execute(w, p)
}

func handleFormSubmit(w http.ResponseWriter, r *http.Request) {

}
