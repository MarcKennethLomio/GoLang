package main

import(
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct{
	Email string `json:"Email"`
	Password string `json:"password"`
}

var Users []User

func handleRequest(){
   initializeRoutes()
}

func initializeRoutes(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage).Methods("GET")
	myRouter.HandleFunc("/users",returnAllUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000",myRouter))
}

func main(){
	initializeUsersMockData()
	handleRequest()
}

func initializeUsersMockData(){
	Users = []User{
		User{
			Email:"marckenneth.lomio@gmail.com",
			Password:"password@123",
		},
		User{
			Email:"melrose.mejidana@gmail.com",
			Password:"password@456",
		},
	}
}

func returnAllUsers(w http.ResponseWriter, r *http.Request){
  json.NewEncoder(w).Encode(Users)
}

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Println("Welcome to homepage")
}