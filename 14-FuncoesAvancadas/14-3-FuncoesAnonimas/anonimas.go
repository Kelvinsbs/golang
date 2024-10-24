package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Passando parametros")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametros")

	fmt.Println(retorno)
}
