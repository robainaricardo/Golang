// This is a test file for the mongoapi.
package mongoapi

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// TestStartAndCloseConnection tests StartConnection and CloseConnection functions.
func TestStartAndCloseConnection(t *testing.T) {
	URI := "mongodb://localhost:27017"
	client := StartConnection(URI)

	if client == nil {
		t.Error("The test fail in StartConnection function.")
	}

	err := CloseConnection(*client)

	if err != nil {
		t.Error("The test fail ain CloseConnection.", err)
	}
}

// TesteClearConnection tests ClearCollection function.
func TestClearCollection(t *testing.T) {
	URI := "mongodb://localhost:27017"
	client := StartConnection(URI)
	collection := client.Database("golang-test").Collection("users-test")

	// Clear all documents of a collection
	ClearCollection(*client, *collection)

	// After clear all collection the result must be zero.
	users, err := QueryUsers(*client, *collection, bson.D{{}})

	if err != nil {
		t.Error("The test fail on QueryUsers inside TestClearCollection function.")
	}

	if len(users) != 0 {
		t.Error("The test fail in TestClearCollection the result must be 0 and was ", len(users))
	}
}

//	TestInsertAndQueryUser tests inserting and querying a user.
func TestInsertAndQueryUser(t *testing.T) {
	URI := "mongodb://localhost:27017"
	client := StartConnection(URI)
	collection := client.Database("golang-test").Collection("users-test")
	ClearCollection(*client, *collection)

	// Testing a missing user.
	query := bson.D{{"name", "Test User"},
		{"email", "testuser@mail.com"},
		{"age", 19}}

	// Testing a absent user
	_, err := QueryUser(*client, *collection, query)

	if err == nil {
		t.Error("The test fail in QueryUser function.")
	}

	// Inserting a user
	newUser := User{"Test User", "testuser@mail.com", 19}
	err = InsertUser(*client, *collection, newUser)

	if err != nil {
		t.Error("Error on InsertUser")
	}

	//	Quering the a existent user
	var user User
	user, err = QueryUser(*client, *collection, query)

	if err != nil {
		t.Error("Error querying a existent user", err)
	}

	if user != newUser {
		t.Error("The users must be the same.", user, newUser)

	}

}

// TestQueryAllUsersAndUpdateAndDelete tests querying all the users, update and delete a user.
func TestQueryAllUsersAndUpdateAndDelete(t *testing.T) {
	URI := "mongodb://localhost:27017"
	client := StartConnection(URI)
	collection := client.Database("golang-test").Collection("users-test")
	ClearCollection(*client, *collection)

	user1 := User{"User 1", "anyemail@mail.com", 1}
	user2 := User{"User 2", "anyemail2@mail.com", 2}
	allUsers := []User{user1, user2}

	err := InsertUser(*client, *collection, user1)
	err = InsertUser(*client, *collection, user2)

	if err != nil {
		t.Error("Error during insert users.")
	}

	var users []User
	users, err = QueryUsers(*client, *collection, bson.D{{}})

	if err != nil {
		t.Error("Erro ao consultar os usuarios")
	}

	if users[0] != allUsers[0] || users[1] != allUsers[1] || len(users) != len(allUsers) {
		t.Error("The users must be the same.", users, allUsers)
	}

	//Update user
	query := bson.D{{"name", users[0].Name}, {"email", users[0].Email}, {"age", users[0].Age}}
	update := bson.D{{"$set", bson.D{{
		"name", "User updated"}}},
		{"$set", bson.D{{
			"age", 100}}}}

	err = UpdateUser(*client, *collection, query, update)

	if err != nil {
		t.Error("Erro ao atualizar o usuário")
	}

	query = bson.D{{"name", "User updated"}, {"email", users[0].Email}, {"age", 100}}
	var updatedUser User
	updatedUser, err = QueryUser(*client, *collection, query)

	if err != nil {
		t.Error("Erro ao consultar o usuario atualizado")
	}

	if users[0] == updatedUser {
		t.Error("The users must be diferent", users[1], updatedUser)
	}

	err = DeleteUser(*client, *collection, query)

	if err != nil {
		t.Error("Erro ao deletar o usuario atualizado")
	}

	updatedUser, err = QueryUser(*client, *collection, query)

	if err == nil {
		t.Error("Se o usuario tivesse sido excluido deveria dar um erro")
	}

}
