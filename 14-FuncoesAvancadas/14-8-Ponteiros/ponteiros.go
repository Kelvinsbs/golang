package main

import "fmt"

func inverterSinal(numero int) int { // passando o valor por copia
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) { // passando o valor por referencia
	*numero = *numero * -1 // * e pra dizer que vai pegar o valor do ponteiro e nao a referencia de memoria
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) // & = falar que esta mandando um ponteiro
	fmt.Println(novoNumero)
}
