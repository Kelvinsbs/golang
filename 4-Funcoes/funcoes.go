package main

import "fmt"

// o padrao da funcao é func (parametro e seu respectivo tipo (caso tenha parametros)) tipo do retorno
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(6, 3)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto da funcao 1")
	fmt.Println(resultado)

	// o _ serve para o go entender que vc nao quer;nao vai usar o segundo parametro
	resultadoSoma, _ := calculosMatematicos(19, 4)
	fmt.Println(resultadoSoma)
}
