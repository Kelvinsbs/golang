package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array e slices")
	// array é uma lista com tamanho fixo, e slice é uma lista sem tamanho fixo

	var array1 [5]string //os itens dentro de um array só pode ter um tipo
	array1[0] = "POsição 1"
	fmt.Println(array1)

	array2 := [5]string{"posicao 1", "posicao 2", "posicao 3", "posicao 4", "posicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICE
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 29, 20}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 21)

	fmt.Println(slice)

	// slice funciona como um ponteiro, ou seja se mudar o parametro do array vai mudar o valor do slice
	slice2 := array2[1:3] // n:n o primeiro é inclusivo (inclui ele na listagem), ja o segundo e exclusivo (pega uma posicao anterior a ele)
	//n:n primeiro n equivalente a >= e o segundo n equivalente a <
	fmt.Println(slice2)

	array2[1] = "posição alterada!"
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	fmt.Println("------------------")
	slice3 := make([]float32, 10, 11) // tipo, tamanho, capacidade
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	fmt.Println("------------------")

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 42)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
