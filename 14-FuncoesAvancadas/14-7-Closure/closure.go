package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"

	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova() //funcao closure, sempre vai ter uma "memoria" de onde ela veio
}
