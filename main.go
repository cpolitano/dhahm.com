package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/listen", listen)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func listen(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "listen.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
