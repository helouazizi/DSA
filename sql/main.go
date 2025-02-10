package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

func New_DB() (*Database, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func main() {
	db, err := New_DB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (username, password)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database created seccusfully")
	http.HandleFunc("/login", Login)
	fmt.Println("http://localhost:8080/login")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed\n"))
		return
	}

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	//fmt.Println(username, password)
	db, err := New_DB()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error\n"))
		return
	}
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?,?)", username, password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error\n"))
		return
	}

	sqlres, err := db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error\n"))
		return
	}
	defer sqlres.Close()
	var user , pass string
	for sqlres.Next() {
		err := sqlres.Scan(&user,&pass)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user)
	}
	w.Write([]byte("hello from server\n"))
}
