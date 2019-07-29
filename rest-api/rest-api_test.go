package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestGETMethod(t *testing.T) {

	out, err := exec.Command("curl", "localhost:8081/users/testuser@gmail.com").Output()

	output := string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, " {name: Test User, email: testuser@gmail.com, age: 100} ") == 0 {
		t.Error("O retorno deveria ser  {name: Test User, email: testuser@gmail.com, age: 100}, e foi: ", output)
	}

}

func TestPOSTMethod(t *testing.T) {

	out, err := exec.Command("curl", "localhost:8081/users", "-X", "POST", "-d", "name=TestUser2&email=testuser2@gmail.com&age=100").Output()

	output := string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, "400 Bad Request") == 0 {
		t.Error("Erro ao inserir usuario", output)
	}

}

func TestDELETEMethod(t *testing.T) {

	out, err := exec.Command("curl", "localhost:8081/users/testuser2@gmail.com", "-X", "DELETE").Output()

	output := string(out[:])
	fmt.Println(output)

	if err != nil {
		t.Error("O comando não foi executado com secesso.", err)
	}

	if strings.Compare(output, "") != 0 {
		t.Error("Não foi possível deletar o usuario")
	}

}
