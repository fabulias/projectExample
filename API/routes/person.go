package routes

import "github.com/gorilla/mux"
import "github.com/fabulias/projectExample/API/controller"
import "net/http"

func RouteringPerson() (http.Handler) {
	router := mux.NewRouter()
	router.HandleFunc("/people", controller.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", controller.GetPerson).Methods("GET")
	router.HandleFunc("/people", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", controller.DeletePerson).Methods("DELETE")
	return router
}
