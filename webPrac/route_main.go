package main

import (
	"fmt"
	"net/http"
)

func errorHandleFunc(writer http.ResponseWriter, request *http.Request) {
	// TODO
	// handle error
	fmt.Println("error~~~~~~~~~~~")
}

func index(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := getSessionIdFromRequest(request)

	if err != nil {
		tempFiles.ExecuteTemplate(writer, "index.html", nil)
	} else {
		getUserFiles(writer, request, sessionId)
	}
	// t := template.Must(template.ParseFiles("./templates/user.html", "./templates/user_body.html", "./templates/user_header.html"))
	// t.ExecuteTemplate(writer, "user.html", testUserFile())
}
