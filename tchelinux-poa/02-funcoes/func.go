package main

import (
	"fmt"
	"os"
)

func main() {
	total := soma(1, 2)
	fmt.Println(total)

	q, r := div(6, 3)
	println("q = ", q, "r = ", r)

	args := os.Args[1:]
	hello(args...)
}

// Soma dois numeros inteiros
func soma(n1 int, n2 int) int {
	return n1 + n2
}

// Divide dois numeros inteiros. Retorna o quociente e o resto dessa operação
func div(nu int, de int) (int, int) {
	q := nu / de
	r := nu % de
	return q, r
}

// Função com parâmetro variádico. Retorna Hello <name> para todos os nomes enviado como parâmetro para o programa
func hello(names ...string) {
	for _, name := range names {
		fmt.Println("Hello", name)
	}
}
