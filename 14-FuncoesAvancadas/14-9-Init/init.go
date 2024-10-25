package main

import "fmt"

var n int

func init() { // SÓ PODE TER UM POR ARQUIVO
	fmt.Println("Executando a funcao init")
	n = 10
}

func main() { // SÓ PODE TER UM POR PACOTE
	fmt.Println("Executando a funcao main")
	fmt.Println(n)
}
