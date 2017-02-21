package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
	"encoding/json"
	"fmt"
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
    params := mux.Vars(req)
    for _, item := range people {
        /*checking if the ID of person entered is mathing the 
        existing ID*/
        if item.ID == params ["id"]{
            fmt.Println("Request for GetPersonEndpoint has been made")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    /*if the id you are looking for does not exist 
    it will just return an ampty struct */
    json.NewEncoder(w).Encode(&Person{})
}



func GetPeopleEndPoint (w http.ResponseWriter, req *http.Request){
    
    json.NewEncoder(w).Encode(people)
}


func CreatePersonEndPoint (w http.ResponseWriter, req *http.Request){
    params := mux.Vars(req)
    var person Person
    _=json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people,person)
    json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint (w http.ResponseWriter, req *http.Request){
    params :=mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"]{
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(people)
}

func main() {
    router := mux.NewRouter()


    /*Sample Set to define people*/
    people = append(people, Person{ID: "1", Firstname: "Valeriia", Lastname:"Savenko", 
        Address: &Address{City: "Brisband", State: "QLD"}})
    people = append(people, Person{ID: "2", Firstname: "Nick", Lastname:"Savenko", })


/*what you will be entering in url for example
---> if you type in localhost:12345/people -> it will get you all people from the array
---> if you type in localhost:12345/people/1 -> it will get you person under id 1*/
    router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndPoint ).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")


    log.Fatal(http.ListenAndServe(":12345", router))

}
