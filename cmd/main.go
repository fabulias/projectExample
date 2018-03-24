package main

import "net/http"
import "log"
import "github.com/fabulias/projectExample/API/routes"

func main() {
	r := routes.RouteringPerson()
	log.Fatal(http.ListenAndServe(":8000", r))
}