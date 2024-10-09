package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 3 - 2
	divisao := 10 / 4
	multiplicacao := 3 * 3
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// no Golang nao e possivel fazer nada com variaveis de tipos diferentes
	// var numero1 int16 = 10 // var numero2 int32 = 25 // soma1 := numero1 + numero2

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma1 := numero1 + numero2

	fmt.Println(soma1)

	// ARITMETICOS FIM

	// ATRIBUICAO
	var variavel1 string = "string"
	variavel2 := "string2"

	fmt.Println(variavel1, variavel2)
	// ATRIBUICAO FIM

	// OPERADORES RELACIONAIS
	// operadores relacionais sempre vao retornar um bool (true ou false)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// OPERADORES RELACIONAIS FIM

	// OPERADORES LOGICOS
	fmt.Println("-------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	// OPERADORES LOGICOS FIM

	// OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15 // mesma coisa que o numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 3 // mesma coisa que o numero = numero * 3
	fmt.Println(numero)

	numero /= 10 // numero = numero / 10
	fmt.Println(numero)

	numero %= 3 // numero = numero % 3
	fmt.Println(numero)
	// OPERADORES UNARIOS FIM

	// OPERADOR TERNARIO //nao existe :C

	// OPERADOR TERNARIO FIM

}
