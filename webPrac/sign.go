package main

import (
	"cache"
	"data"
	"fmt"
	"log"
	"net/http"
)

func signInFunc(writer http.ResponseWriter, request *http.Request) {
	tempFiles.ExecuteTemplate(writer, "sign_in.html", nil)
}

func signUpFunc(writer http.ResponseWriter, request *http.Request) {
	tempFiles.ExecuteTemplate(writer, "sign_up.html", nil)
}

func doSignUp(writer http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("username")
	password := request.URL.Query().Get("password")

	if len(username) == 0 || len(password) == 0 {
		writer.Write([]byte("username or password is null"))
		return
	}

	u := data.NewUser(username, password)
	err := data.InsertUser(u)

	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(writer, "success")
}

func doSignIn(writer http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("username")
	password := request.URL.Query().Get("password")

	if len(username) == 0 || len(password) == 0 {
		writer.Write([]byte("username or password is null"))
		return
	}

	u, err := data.GetUserByUserName(username)
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}

	if data.Encode(nil, password) == u.PassWord {
		ck, err := generateCookieFromUser(u)
		if err != nil {
			log.Println(err)
			writer.Write([]byte(err.Error()))
		} else {
			http.SetCookie(writer, ck)
			writer.Write([]byte("success"))
			cache.Put(u.Id, u)
		}
	} else {
		writer.Write([]byte("Incorrect username or password"))
	}
}
