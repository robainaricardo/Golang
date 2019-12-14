package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// Defer
	file, _ := os.Open("file01.txt")
	fmt.Println("Abrir o arquivo", file)
	fmt.Println("Manipular o arquivo", file)
	fmt.Println("Fechar o arquivo", file)

	// Concorrência
	r1 := 0
	go etapa1(&r1)
	go etapa2(&r1)
	go etapa3(&r1)

	time.Sleep(2 * time.Second)

	println(r1)
	fmt.Println("done")
}

// Função que representa a primeira etapa de um algoritmo hipotético
func etapa1(valor *int) {
	fmt.Println("Etapa 1")
	*valor++
}

// Função que representa a segunda etapa de um algoritmo hipotético
func etapa2(valor *int) {
	fmt.Println("Etapa 2")
	*valor *= (*valor)
}


// Função que representa a terceira etapa de um algoritmo hipotético
func etapa3(valor *int) {
	fmt.Println("Etapa 3")
	*valor += 1025
}
