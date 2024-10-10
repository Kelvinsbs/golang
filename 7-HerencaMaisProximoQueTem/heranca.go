package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

// type estudante struct {
// 	pessoa pessoa // se passar dessa forma, para acessar posteriormente teria que acessar estudante.pessoa.nome etc
// 	curso     string
// 	faculdade string
// }

func main() {
	fmt.Println("heranca")

	p1 := pessoa{"kleber", "machado", 21, 174}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
}
