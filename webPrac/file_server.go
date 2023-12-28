package main

import (
	"cache"
	"data"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type UserFileDao struct {
	UserName    string
	FilePath    []string
	FilePathLen int
	FileInfos   []FileInfo
	FIleInfoLen int
}

type FileInfo struct {
	Name     string
	BasePath string
	IsDir    bool
}

func getUserFiles(writer http.ResponseWriter, request *http.Request, sessionId string) {
	u, ok := cache.Get(sessionId)

	if !ok {
		errMsg := fmt.Sprintf("user not log in ! info[package : main, file : file_server.go, func : getUserFiles], args[argName : sessionId, argValue : %s]\n", sessionId)
		log.Println(errMsg)
		fmt.Fprint(writer, errMsg)
	}

	userInfo := u.(*data.User)

	reqPath := request.URL.Query().Get("path")
	basePath := getUserBaseDir(userInfo)

	if !fileExists(basePath) {
		os.Mkdir(basePath, os.ModePerm)
	}

	dao := UserFileDao{
		UserName: userInfo.UserName,
	}

	filePath := []string{}

	fileInfos := []FileInfo{}

	// if there is no request path, return the base path
	if reqPath == "" || reqPath == "Home" {
		filePath = append(filePath, basePath)
		entries, err := os.ReadDir(basePath)
		if err != nil {
			errMsg := fmt.Sprintf("read base path error! info[package : main, file : file_server.go, func : getUserFiles], args[argName : basePath, argValue : %s]\n", basePath)
			log.Println(errMsg)
			log.Println(err)
			fmt.Fprint(writer, errMsg)
		} else {
			for _, entry := range entries {
				fileInfos = append(fileInfos, FileInfo{
					Name:     entry.Name(),
					BasePath: basePath,
					IsDir:    entry.Type().IsDir(),
				})
			}
		}

	} else {
		reqPath = strings.Replace(reqPath, "Home", userInfo.Id, 1)
		reqPath = strings.Join([]string{"./", FileBaseDir, reqPath}, "")

		info, err := os.Stat(reqPath)
		if err != nil {
			// TODO
		}

		isdDir := info.IsDir()
		if isdDir {

			entries, err := os.ReadDir(reqPath)
			if err != nil {

			}

			for _, entry := range entries {
				entry.Name()
			}
		}
	}

	dao.FileInfos = fileInfos
	dao.FilePath = filePath

	tempFiles.ExecuteTemplate(writer, "user.html", dao)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func testUserFile() UserFileDao {
	u := UserFileDao{}

	u.UserName = "Aresus"

	u.FilePath = []string{"Home", "123", "321"}

	u.FileInfos = []FileInfo{
		FileInfo{Name: "test1", BasePath: "Home/123/321", IsDir: true},
		FileInfo{Name: "test2", BasePath: "Home/123/321", IsDir: true},
		FileInfo{Name: "test3", BasePath: "Home/123/321", IsDir: false},
	}

	u.FilePathLen = 2
	u.FIleInfoLen = 2

	return u
}
