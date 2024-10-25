package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() { // SÓ PODE TER UM POR PACOTE
	fmt.Println("Executando a funcao main")
	usuario1 := usuario{"usuario um", 21}
	usuario1.salvar()

	usuario2 := usuario{"davi", 17}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

	maiorDeIdade2 := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade2)
}
