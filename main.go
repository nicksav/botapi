package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id, omitempty"`
	Firstname string   `json:"firstname, omitempty"`
	Lastname  string   `json:"lastname, omitempty"`
	Address   *Address `json:"address, omitempty"`
}

type Address struct {
	City  string `json:"city, omitempty"`
	State string `json:"state, omitempty"`
}

var people []Person

/*GetPersonEndPoint -- this is the same as GET -> which is used to retieve resourse based on specified
ID */
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		/*checking if the ID of person entered is mathing the
		  existing ID*/
		if item.ID == params["id"] {
			fmt.Println("Request for GetPersonEndpoint has been made")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	/*if the id you are looking for does not exist
	  it will just return an ampty struct */
	json.NewEncoder(w).Encode(&Person{})
}

/*GetPeopleEndPoint -- this is the same as GET -> which is used to retieve resourse */
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {

	json.NewEncoder(w).Encode(people)
}

/*CreatePersonEndPoint -- this is the same as POST -> which is used to create resourse*/
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	/*adding newly created person to the struct*/
	people = append(people, person)
	/*returning back to the page to view*/
	json.NewEncoder(w).Encode(people)
}

/*DeletePersonEndPoint -- this is DELETE -> which is used to delete resourse*/
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	/*retieve the parameter entered by the user*/
	params := mux.Vars(req)

	/*cheking if the param "ID" entered by the user is matching any
	  existing ID*/
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

/*AddPersonLastNameEndPoint -- adds person's last name by ID*/
func AddPersonLastNameEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	person.Lastname = params["lastname"]
	/*adding newly created person to the struct*/
	people = append(people, person)
	/*returning back to the page to view*/
	json.NewEncoder(w).Encode(people)

}

func main() {
	router := mux.NewRouter()

	/*Sample Set to define people*/
	people = append(people, Person{ID: "1", Firstname: "Valeriia", Lastname: "Savenko",
		Address: &Address{City: "Brisband", State: "QLD"}})
	people = append(people, Person{ID: "2", Firstname: "Nick", Lastname: "Savenko"})

	/*what you will be entering in url for example
	  ---> if you type in localhost:12345/people -> it will get you all people from the array
	  ---> if you type in localhost:12345/people/1 -> it will get you person under id 1*/

	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	router.HandleFunc("/people/{id}/{lastname}", AddPersonLastNameEndPoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}
