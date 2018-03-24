package controller 

import "net/http"
import  "log"
//import "github.com/gorilla/mux"
import "github.com/fabulias/projectExample/API/schema"
import "github.com/fabulias/projectExample/API/model"
import "encoding/json"

func GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("Todo ok :)")

}
func GetPerson(w http.ResponseWriter, r *http.Request) {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person schema.Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    log.Println("In controller -> ", person)
    model.CreatePerson(person)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}