package main

import(
	"net/http"
	_ "docs"
    services "services"
    controllers "controllers"
)


type UserDTO struct{
        Email string    `json:"email"` 
        Password string `json:"password"`
        }
        

// @title Go Rest API
// @version 1.0
// @description Go Rest API using Swaggo & SQL SERVER DB
// @contact.name Marc Kenneth Lomio
// @contact.email marckenneth.lomio@gmail.com
// @host localhost:9000
// @BasePath /
func main() { 
	controllers.HandleRequests()
}

// @Summary Migrate tables to the sql server
// @Description
// @Tags migration
// @Produce json
// @Success 200
// @Router /migration [post]
func CreateDatabaseSchema(w http.ResponseWriter, r *http.Request){
    services.CreateDatabaseSchema(w,r)
   }

// @Summary Create User
// @Description
// @Tags users
// @Produce json
// @Param user body UserDTO true "create user dto"
// @Success 200 {object} UserDTO
// @Router /user [post]
func CreateNewUser(w http.ResponseWriter, r *http.Request){
    services.CreateNewUser(w,r)
   }

// @Summary Update User
// @Description
// @Tags users
// @Produce json
// @Param user body UserDTO true "update user dto"
// @Success 200 {object} UserDTO
// @Router /user [put]
func UpdateUser(w http.ResponseWriter, r *http.Request){
    services.UpdateUser(w,r)
   }

// @Summary Get All users
// @Description
// @Tags users
// @Produce json
// @Success 200 
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request){
   services.GetAllUsers(w,r)
}