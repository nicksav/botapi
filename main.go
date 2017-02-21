package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
	"encoding/json"
)

type Person struct {
    ID   string `json:"id, omitempty"`
    Firstname   string `json:"firstname, omitempty"`
    Lastname   string `json:"lastname, omitempty"`
    Address *Address`json:"address, omitempty"`

}

type  Address struct {
    City string `json:"city, omitempty"`
    State   string `json:"state, omitempty"`

}

var people []Person

func GetPersonEndPoint (w http.ResponseWriter, req *http.Request){
    
}

func GetPeopleEndPoint (w http.ResponseWriter, req *http.Request){
    
    json.NewEncoder(w).Encode(people)
}

func CreatePersonEndPoint (w http.ResponseWriter, req *http.Request){
    
}

func DeletePersonEndPoint (w http.ResponseWriter, req *http.Request){
    
}

func main() {
    router := mux.NewRouter()
    /*Sample Set to define people*/
    people = append(people, Person{ID: "1", Firstname: "Valeri", Lastname:"Savenko", 
        Address: &Address{City: "Brisband", State: "QLD"}})

    people = append(people, Person{ID: "2", Firstname: "Nick", Lastname:"Savenko", })

    router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPeopleEndPoint).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")


    log.Fatal(http.ListenAndServe(":12345", router))

}
