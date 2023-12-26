package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	host     string
	password string
	db       *sql.DB
)

func init() {
	var err error
	host = "127.0.0.1"
	password = "root"
	db, err = sql.Open("mysql", "root:root@/aresus_cloud")
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(6)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Database connection has established")
}

func GetUserByUserName(username string) (*User, error) {

	if len(username) == 0 {
		errMsg := fmt.Sprintf("Illegal username error, info[package : data, file : mysql_utils.go, func : GetUserByUserName], args[argName : username, argValue : %s]", username)
		return nil, errors.New(errMsg)
	}

	rows, err := db.Query("select * from sys_user where user_name=?;", username)

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		u := &User{}
		rows.Scan(&u.Id, &u.UserName, &u.PassWord, &u.Role)
		return u, nil
	}

	errMsg := fmt.Sprintf("No user found error, info[package : data, file : mysql_utils.go, func : GetUserByUserName], args[argName : username, argValue : %s]", username)

	return nil, errors.New(errMsg)
}

func InsertUser(u *User) error {
	if u == nil {
		return errors.New("User is nil, info[package : data, file : mysql_utils.go, func : InsertUser]")
	}

	_, err := db.Exec("insert into sys_user (user_id, user_name, pass_word, role) values (?, ?, ?, ?)", u.Id, u.UserName, u.PassWord, u.Role)

	if err != nil {
		return err
	}

	return nil
}

func Test() *User {
	u := &User{}
	db.QueryRow("select * from sys_user where user_name = ? and pass_word = ?;", "admin", "admin").Scan(&u.Id, &u.UserName, &u.PassWord, &u.Role)
	return u
}

func TestInsert() {
	_, err := db.Exec("insert into sys_user (user_id, user_name, pass_word, role) values (?, ? , ?, ?)", "test", "test", "test", "test")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Insert test success!")

}
