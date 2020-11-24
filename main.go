package main

// import of all dependencies
// github.com/gorilla/mux gives us a minimal server and router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Types
type superheroe struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Alias string `json:"Alias"`
}

type superheroes []superheroe

// Persistence
var superheroesList = superheroes{
	{Name: "Thor", Alias: "Thor Odinson", ID: 1},
	{Name: "Batman", Alias: "Bruce Wayne", ID: 2},
	{Name: "Iron Man", Alias: "Tony Stark", ID: 3},
	{Name: "Superman", Alias: "Clark Kent", ID: 4},
}

func ping(res http.ResponseWriter, req *http.Request) {
	// we return a simple welcome to my GO API message back to the client
	fmt.Fprintf(res, "Welcome the SuperHero GO API")
}

func addSuperHeroe(res http.ResponseWriter, req *http.Request) {
	var newHeroe superheroe // we set up a newHeroe variable that is of type "superheroe"

	/*
		we use ioutil.ReadAll to read all the data from the request body (to extract the data)
		this returns a value that we save on reqBody variable or an error
	*/

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(res, "Insert a Valid SuperHeroe Data")
	}

	/*
		we transform the reqBody variable to a valid json and assign that json to the
		newTask variable
	*/

	json.Unmarshal(reqBody, &newHeroe)
	fmt.Println("newHeroe")
	newHeroe.ID = len(superheroesList) + 1

	/*
		we append the new task to our array of tasks
	*/
	superheroesList = append(superheroesList, newHeroe)

	/*
		we set the headers for the response back to the client
		and then using the jsonNewEncoder function we encode the array of task
		back to the client
	*/
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newHeroe)
}

func getSuperheroes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(superheroesList)
}

func getSuperheroe(res http.ResponseWriter, req *http.Request) {
	/*
		we extract the query variables from the req object (/{id})
	*/
	vars := mux.Vars(req)

	/*
		with strconv.Atoi we parse a string value to a number value
		superheroeID, err := strconv.Atoi(vars["id"])
	*/
	superheroeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	for _, heroe := range superheroesList {
		/*
			in this for loop we check if a heroe id match the superheroeID extracted from the req object
		*/
		if heroe.ID == superheroeID {
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(heroe)
		}
	}
}

func updateSuperheroe(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	superheroeID, err := strconv.Atoi(vars["id"])
	var updatedTask superheroe

	if err != nil {
		fmt.Fprintf(res, "Invalid ID")
	}
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(res, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for index, heroe := range superheroesList {
		if heroe.ID == superheroeID {
			/*
				if we find the superheroe we need to update his information
				we remove the heroe found in the superheroesList with the append function
				for this we need to append all the super heroes except the current one
				for that we said to append that append the heroes before the one with index
				and all the others after the one with the current index
			*/
			superheroesList = append(superheroesList[:index], superheroesList[index+1:]...)
			updatedTask.ID = superheroeID // we set up the id for the current heroe
			superheroesList = append(superheroesList, updatedTask)
			fmt.Fprintf(res, "The task with ID %v has been updated successfully", superheroeID)
		}
	}
}

func deleteSuperheroe(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	superheroeID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(res, "Invalid User ID")
		return
	}

	for index, heroe := range superheroesList {
		if heroe.ID == superheroeID {
			superheroesList = append(superheroesList[:index], superheroesList[index+1:]...)
			fmt.Fprintf(res, "The task with ID %v has been remove successfully", superheroeID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", ping)
	router.HandleFunc("/superheroes", addSuperHeroe).Methods("POST")
	router.HandleFunc("/superheroes", getSuperheroes).Methods("GET")
	router.HandleFunc("/superheroes/{id}", getSuperheroe).Methods("GET")
	router.HandleFunc("/superheroes/{id}", deleteSuperheroe).Methods("DELETE")
	router.HandleFunc("/superheroes/{id}", updateSuperheroe).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5000", router))
}
