package services

import(
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	repositories "repositories"
	entities "entities"
)

func CreateDatabaseSchema(w http.ResponseWriter, r *http.Request){
  fmt.Println("services CreateDatabaseSchema")
  repositories.SchemaMigration()
  w.WriteHeader(http.StatusCreated)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request){
  fmt.Println("services createnewUser")

  reqBody, _ := ioutil.ReadAll(r.Body)
  var user entities.User
  json.Unmarshal(reqBody,&user)
  user = repositories.CreateNewUser(user,user.Password,false)
  json.NewEncoder(w).Encode(user)
  w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
  fmt.Println("services updateuser")
  reqBody, _ := ioutil.ReadAll(r.Body)
  var user entities.User
  json.Unmarshal(reqBody,&user)
  repositories.UpdateUser(user)
  json.NewEncoder(w).Encode(user)
  w.WriteHeader(http.StatusCreated)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
  fmt.Println("services getAllusers")
  var users []entities.User
  users = repositories.GetAllUsers()
  json.NewEncoder(w).Encode(users)
  w.WriteHeader(http.StatusOK)
}