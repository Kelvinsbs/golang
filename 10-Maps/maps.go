package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// sempre vai ser o mesmo tipo
	usuario := map[string]string{ // dentro do conchetes e o tipo da chave, fora do conchetes e o tipo do valor
		"nome":      "Pedro",
		"Sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joão",
			"segundo":  "kleber",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	// deletar uma chave
	delete(usuario2, "nome")

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "capricornio",
	}

	fmt.Println(usuario2)

}
