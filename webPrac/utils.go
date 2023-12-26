package main

import (
	"data"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"
)

const cookie_key = "_aresus_cookie"

func getSessionIdFromRequest(request *http.Request) (string, error) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func readAllTemplates(filePackage string) *template.Template {

	entries, err := os.ReadDir(filePackage)

	if err != nil {
		panic(err)
	}

	filePaths := []string{}
	for _, entry := range entries {
		fileName := entry.Name()
		if strings.Contains(fileName, "html") {
			filePaths = append(filePaths, filePackage+string(os.PathSeparator)+fileName)
		}
	}

	return template.Must(template.ParseFiles(filePaths...))
}

func generateCookieFromUser(u *data.User) (*http.Cookie, error) {
	if u == nil {
		errMsg := fmt.Sprintf("user info is null, info[package : main, file : uutils.go, func : generateCookieFromUser], args[argName : u, argValue : %v]", u)
		return nil, errors.New(errMsg)
	}

	return &http.Cookie{
		Name:     cookie_key,
		Value:    u.Id,
		HttpOnly: true,
		Expires:  time.Now().Add(30 * time.Minute),
	}, nil
}

func getUserBaseDir(u *data.User) string {
	userId := u.Id
	return strings.Join([]string{FileBaseDir, userId}, "")
}
