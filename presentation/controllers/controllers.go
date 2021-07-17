package controllers

import(
	"github.com/gorilla/mux"
	"log"
	"net/http"
	services "services"
	httpSwagger "github.com/swaggo/http-swagger"
)

func HandleRequests() {
	InitializeRoutes()
}

func InitializeRoutes(){
	InitRoutesByGorillaMux()
}

func InitRoutesByGorillaMux(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/migration",services.CreateDatabaseSchema).Methods("POST")
	myRouter.HandleFunc("/user", services.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user",services.CreateNewUser).Methods("POST")
	myRouter.HandleFunc("/users",services.GetAllUsers).Methods("GET")
	myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
    log.Fatal(http.ListenAndServe(":9000",myRouter))
}