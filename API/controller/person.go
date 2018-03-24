package controller 

import "net/http"
import  "log"

func GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("Todo ok :)")
}
func GetPerson(w http.ResponseWriter, r *http.Request) {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}