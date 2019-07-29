package main

import (
	mong "Golang/mongo-api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {

	URI := "mongodb://localhost:27017"
	client := mong.StartConnection(URI)
	collection := client.Database("golang-test").Collection("users")

	defer mong.CloseConnection(*client)

	var users []mong.User

	users = mong.QueryUsers(*client, *collection, bson.D{{}})

	for _, u := range users {
		fmt.Println(u)
	}

	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	email := vars["email"]

	URI := "mongodb://localhost:27017"
	client := mong.StartConnection(URI)
	collection := client.Database("golang-test").Collection("users")

	defer mong.CloseConnection(*client)

	consulta := bson.D{{"email", email}}
	user := mong.QueryUser(*client, *collection, consulta)

	json.NewEncoder(w).Encode(user)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("name")
	userMail := r.FormValue("email")
	userAge := r.FormValue("age")
	fmt.Println(userName, userMail, userAge)
	age, err := strconv.Atoi(userAge)
	if err != nil {
		log.Fatal(err)
	}
	user := mong.User{userName, userMail, age}
	fmt.Println(user)

	URI := "mongodb://localhost:27017"
	client := mong.StartConnection(URI)
	collection := client.Database("golang-test").Collection("users")

	defer mong.CloseConnection(*client)

	mong.InsertUser(*client, *collection, user)

}

func putUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	email := vars["email"]

	//	Dados para a atualização
	userName := r.FormValue("name")
	userMail := r.FormValue("email")
	userAge := r.FormValue("age")

	age, err := strconv.Atoi(userAge)
	if err != nil {
		log.Fatal(err)
	}
	user := mong.User{userName, userMail, age}

	URI := "mongodb://localhost:27017"
	client := mong.StartConnection(URI)
	collection := client.Database("golang-test").Collection("users")
	defer mong.CloseConnection(*client)

	consulta := bson.D{{"email", email}}
	atualizacao := bson.D{
		{"$set", bson.D{{"name", user.Name}}},
		{"$set", bson.D{{"email", user.Email}}},
		{"$set", bson.D{{"age", user.Age}}}}
	mong.UpdateUser(*client, *collection, consulta, atualizacao)

}

//	Delete a user by an email passed in lolcalhost:8081/users/email
func deleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	email := vars["email"]

	URI := "mongodb://localhost:27017"
	client := mong.StartConnection(URI)
	collection := client.Database("golang-test").Collection("users")

	defer mong.CloseConnection(*client)

	consulta := bson.D{{"email", email}}
	mong.DeleteUser(*client, *collection, consulta)

}

// homePage returns a simple message of this API
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a simple REST API writen in Golang that uses MongoDB database.")
}

//	Trata das requisições (mapeia a requisição para a função adequada)
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", getAllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{email}", getUser).Methods("GET")
	myRouter.HandleFunc("/users", postUser).Methods("POST")
	myRouter.HandleFunc("/users/{email}", putUser).Methods("PUT")
	myRouter.HandleFunc("/users/{email}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

//	Função Principal do programa
func main() {
	fmt.Println("API: on")
	defer fmt.Println("API: off")
	handleRequests()

}
