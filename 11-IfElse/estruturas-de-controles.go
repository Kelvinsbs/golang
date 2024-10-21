package main

import "fmt"

func main() {
	fmt.Println("estruturas de controle")

	numero := -10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// esse novo tipo de if, serve para nao precisar criar a variavel em uma linha e a condição/if em outra
	if outroNumero := numero; outroNumero > 0 { //if init
		fmt.Println("Numero é maior do que 0")
	} else if numero < -10 {
		fmt.Println("Numero entre 0 e -10")
	} else {
		fmt.Println("Numero e menor que -10")
	}
	// nao da pra usar a variavel do if init fora do if que ela foi criada

}
