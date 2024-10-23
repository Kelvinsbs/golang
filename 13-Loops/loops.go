package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando I")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"kelvin", "thiago", "gustavo"}

	// key e value do foreach
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// quando nao quiser/preisar do key, so usar o _
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra, string(letra))
	}

	usuario := map[string]string{
		"nome":      "leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// NAO E POSSIVEL FAZER UM FOR PARA UM STRUCT

	// LOOP INFINITO

	for {
		fmt.Println("executando infinitamente")
		time.Sleep(time.Second)
	}
}
