package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS

	// int8, int16, int32, int64 // a quantidade de numeros é a quantidade de bites que esse int vai suportar]
	// int <- variavel sem o valor definido ele vai usar a arquitetura do computador como base (w32 = int32 ou w64 = int64)
	numero := -1000000000
	fmt.Println(numero)

	// uint8, uint16, uint32, uint64 // mesma coisa que o int, porém nao aceita valores negativos
	// uint //unsygned int
	var numero2 uint32 = 100000
	fmt.Println(numero2)

	// alias
	var numero3 rune = 12345 // isso aqui é a mesma coisa que o int32 // int32 = rune
	fmt.Println(numero3)

	var numero4 byte = 123 // isso aqui é a mesma coisa que um uint8 // byte = uint8
	fmt.Println(numero4)
	// NUMEROS INTEIROS FIM

	// NUMEROS REAIS
	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12347346236.45
	fmt.Println(numeroReal2)

	numeroReal3 := 245362.64
	fmt.Println(numeroReal3)
	// NUMEROS REAIS FIM

	// STRING
	// em golang, SEMPRE utilizar aspas duplas pra string
	var str string = "lorem ipsum"
	fmt.Println(str)

	str2 := "dolor sit amet"
	fmt.Println(str2)

	// se colocar um caractere entre aspas simples """seria o equivalente ao char""" ele vai pegar a posicao desse caractere na tabela ASCII
	char := 'B'
	fmt.Println(char)
	// STRING FIM

	var texto int16 // valor inicial do int vai ser 0 e da string vai ser ""
	fmt.Println(texto)

	// BOOLEAN
	var booleano1 bool = true //true ou false
	fmt.Println(booleano1)

	var booleano2 bool //valor inicial do tipo bool e sempre false
	// var booleano2 bool = false // isso aqui é a mesma coisa que a linha de cima
	fmt.Println(booleano2)
	// BOOLEAN FIM

	// ERROR
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
	// ERROR FIM

}
