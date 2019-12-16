package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func inint() {
	tpl = template.Must(template.ParseGlob("template/*"))
}
func idx(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "indexx.html", nil)
	if err != nil {
		log.Println(err)
	}
}
func abt(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
func cnt(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
func aply(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abt)
	http.HandleFunc("/contact", cnt)
	http.HandleFunc("/apply", aply)
	http.ListenAndServe(":8080", nil)
}
