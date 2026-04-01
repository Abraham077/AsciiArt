package server

import (
	"ascii-art-web/internal/ascii"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const statDir = "static"

func RegisterRoutes() {
	fs := http.FileServer(http.Dir(statDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("ascii-art", AsciiHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		WriteErrorPage(w, http.StatusNotFound, "Page not found")
		return
	}
	if r.Method != http.MethodGet {
		WriteErrorPage(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	renderIndex(w, ascii.PageData{})
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteErrorPage(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	if err := r.ParseForm(); err != nil {
		WriteErrorPage(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	text := strings.TrimSpace(r.FormValue("text"))
	banner := r.FormValue("banner")

	if text == "" {
		WriteErrorPage(w, http.StatusBadRequest, "Text cannot be empty")
		return
	}

	if !ascii.IsValidBanner(banner) {
		WriteErrorPage(w, http.StatusBadRequest, "Invalid banner selector")
		return
	}

	result, err := ascii.CreateAscii(text, banner)
	if err != nil {
		if os.IsNotExist(err) {
			WriteErrorPage(w, http.StatusInternalServerError, "Banner file not found")
			return
		}
		if err == os.ErrInvalid {
			WriteErrorPage(w, http.StatusBadRequest, "Unsupported chars in text")
			return
		}
		log.Printf("CreateASCII error: %v", err)
		WriteErrorPage(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	data := ascii.PageData{
		Text:   text,
		Banner: banner,
		Result: result,
	}
	renderIndex(w, data)
}
func renderIndex(w http.ResponseWriter, data ascii.PageData) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles(templateDir + "/index.html")
	if err != nil {
		log.Printf("template index error: %v", err)
		WriteErrorPage(w, http.StatusInternalServerError, "Template error")
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("execute index error: %v", err)
		WriteErrorPage(w, http.StatusInternalServerError, "Template Rendering error")
		return
	}
}
