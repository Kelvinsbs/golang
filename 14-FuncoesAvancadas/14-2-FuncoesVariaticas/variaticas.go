package main

import "fmt"

// vai receber de 1 a n numeros
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

// SÓ PODE TER NO MAX 2 TIPOS DE VALORES
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	// se nao passar nada ou sja vazio, ele vai retornar 0
	totalSoma := soma(1, 2, 4, 5, 6, 123, 5435, 25436)

	fmt.Println(totalSoma)

	escrever("Olá Mundo", 4, 5, 7, 345, 64)
}
