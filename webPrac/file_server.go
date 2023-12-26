package main

import (
	"cache"
	"data"
	"fmt"
	"net/http"
	"os"
)

type UserFileDao struct {
	username  string
	fileInfos []FileInfo
}

type FileInfo struct {
	name     string
	basePath string
	isDir    bool
}

func getUserFiles(writer http.ResponseWriter, request *http.Request, sessionId string) {
	u, ok := cache.Get(sessionId)

	if !ok {
		// TODO
		// unlogin
	}

	userInfo := u.(*data.User)

	reqPath := request.URL.Query().Get("path")
	basePath := getUserBaseDir(userInfo)

	if reqPath != "" {
		getRequestPath(writer, basePath, reqPath, userInfo)
	} else {
		getBaseDir(writer, basePath, userInfo)
	}

}

func getBaseDir(writer http.ResponseWriter, basePath string, userInfo *data.User) {
	if !FileExists(basePath) {
		os.Mkdir(tempFiles.Tree.ParseName, os.ModePerm)
	}

	dao := UserFileDao{
		username:  userInfo.UserName,
		fileInfos: []FileInfo{},
	}

	fmt.Println(dao)
	// TODO
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func getRequestPath(writer http.ResponseWriter, basePath string, reqPath string, userInfo *data.User) {
	// TODO
}
