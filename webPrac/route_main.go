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
	// sessionId, err := getSessionIdFromRequest(request)

	// if err != nil {
	// 	tempFiles.ExecuteTemplate(writer, "index.html", nil)
	// } else {
	// 	getUserFiles(writer, request, sessionId)
	// }
	tempFiles.ExecuteTemplate(writer, "user.html", nil)
}
