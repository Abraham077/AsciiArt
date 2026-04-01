package server

import (
	"ascii-art-web/internal/ascii"
	"html/template"
	"log"
	"net/http"
)

const templateDir = "templates"

func WriteErrorPage(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	tmpl, err := template.ParseFiles(templateDir + "err/.html")
	if err != nil {
		log.Printf("template error %v", err)
		http.Error(w, msg, status)
		return
	}
	data := ascii.ErrorPageData{
		Status:  status,
		Message: msg,
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("execute error %v", err)
		http.Error(w, msg, status)
	}
}
