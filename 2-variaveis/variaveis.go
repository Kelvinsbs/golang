package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1" // declaracao de variavel explicita
	variavel2 := "Variavel 2"           // declaracao de variavel implicita // o nome disso é "inferencia de tipo"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	) // declarar mais de uma variavel ao mesmo tempo de forma explicita

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6" // declarar mais de uma variavel ao mesmo tempo por inferencia de tipo

	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1" // as formas de declarar uma constante é igual as de uma variavel
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
