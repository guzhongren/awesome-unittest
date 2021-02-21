package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

// CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	checkError(err)
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	db.Create(&User{Name: name, Email: email})
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	io.WriteString(w, `{"status": "success"}`)
}

// GetAllUsers 获取所有用户
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	checkError(err)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	var userList []User
	db.Find(&userList)
	response := &Response{
		Status: true,
		Data:   userList,
	}
	byte, err := json.Marshal(response)
	checkError(err)
	io.WriteString(w, string(byte))
}

// HandleHello is root handler
func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	io.WriteString(w, `{ "status": "expected service response"}`)
}

// HandleRequests is server root
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/test", HandleHello).Methods("GET")
	myRouter.HandleFunc("/users", GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func checkError(err error) {
	if err != nil {
		log.Error(`Error:`, err)
	}
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	checkError(err)
	db.AutoMigrate(&User{})
}

func main() {
	initialMigration()
	HandleRequests()
}
