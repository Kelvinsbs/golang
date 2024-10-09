package main

import "fmt"

// de forma resumida o struct é uma "colecao de campos"
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("arquivo structs")

	var u usuario
	u.nome = "Kelvin"
	u.idade = 25
	fmt.Println(u)

	enderecoExemplo := endereco{"rua dos boobos", 2}

	usuario2 := usuario{"golias", 43, enderecoExemplo}
	fmt.Println(usuario2)

	// quando nao tiver todos os dados, só passar qual coluna esta se referindo
	usuario3 := usuario{nome: "alberto"}
	fmt.Println(usuario3)
}
