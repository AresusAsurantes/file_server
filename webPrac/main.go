package main

import (
	"html/template"
	"net/http"
)

const (
	FileBaseDir = "./data"

	TemplateBaseDir = "./templates"
)

var tempFiles *template.Template

func main() {

	tempFiles = readAllTemplates(TemplateBaseDir)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", errorHandleFunc)

	mux.HandleFunc("/signIn", signInFunc)
	mux.HandleFunc("/doSignIn", doSignIn)

	mux.HandleFunc("/signUp", signUpFunc)
	mux.HandleFunc("/doSignUp", doSignUp)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
