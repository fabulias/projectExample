package model

import "github.com/fabulias/projectExample/API/schema"
import "log"

func CreatePerson(person schema.Person) {
	//Insert in db
	log.Println("In model -> ", person) 	
} 