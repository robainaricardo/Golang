//	This is a test file to rest-api.go
//	I use the linux command CURL to consume the API

//IDEIA: Criar um pacote de testes automatizados para api usando o CURL
package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// TestGETMethod
func TestGETMethod(t *testing.T) {

	//	 Testing a existent user
	out, err := exec.Command("curl", "localhost:8081/users/testuser@gmail.com").Output()

	output := string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, " {name: Test User, email: testuser@gmail.com, age: 100} ") == 0 {
		t.Error("O retorno deveria ser  {name: Test User, email: testuser@gmail.com, age: 100}, e foi: ", output)
	}

	//	Testing a non-existent user
	//	 Testing a existent user
	out, err = exec.Command("curl", "localhost:8081/users/nonexistentuser@mail.com").Output()

	output = string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, "404 - Not Found") != 0 {
		t.Error("Should return a Error: 404 - Not Found, and return", output)
	}

	//	Testing get all users
	out, err = exec.Command("curl", "localhost:8081/users").Output()

	output = string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, "400 - Bad Request") == 0 {
		t.Error("Shouldn't return a Error: 400 - Bad Request, and return", output)
	}

	if strings.Compare(output, "[{}]") == 0 {
		t.Error("This should return all the users and has returned: ", output)
	}

}

//	TestPOSTMethod
func TestPOSTMethod(t *testing.T) {

	out, err := exec.Command("curl", "localhost:8081/users", "-X", "POST", "-d", "name=TestUser2&email=testuser2@gmail.com&age=100").Output()

	output := string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, "400 - Bad Request") == 0 {
		t.Error("Erro ao inserir usuario", output)
	}

}

func TestPUTMethod(t *testing.T) {

	out, err := exec.Command("curl", "localhost:8081/users/testuser2@gmail.com", "-X", "PUT", "-d", "name=TestUser3&email=testuser3@gmail.com&age=1003").Output()

	output := string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, "400 - Bad Request") == 0 {
		t.Error("The result shoud be 200 - Ok and was: ", out)
	}

	out, err = exec.Command("curl", "localhost:8081/users/testuser@gmail.com").Output()

	output = string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, " {name: TestUser3, email: testuser3@gmail.com, age: 1003} ") == 0 {
		t.Error("The user wasn't update in the right way", output)
	}

}

func TestDELETEMethod(t *testing.T) {

	out, err := exec.Command("curl", "localhost:8081/users/testuser2@gmail.com", "-X", "DELETE").Output()

	output := string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, "200 - Ok") != 0 {
		t.Error("The result shoud be 200 - Ok and was: ", out)
	}

}
