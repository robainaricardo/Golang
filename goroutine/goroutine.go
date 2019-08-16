package main

import "fmt"

func main() {
	fmt.Println("Começou!")
	go hello("leo")
	go hello("ricardo")
	fmt.Println("Terminou!")

	var input string
	fmt.Scanln(&input)
}

func hello(name string) {
	fmt.Println("Olá ", name)
}
