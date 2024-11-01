package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo!", canal)
	fmt.Println("depois da função escrever começar a ser executada")
	for {
		mensagem, aberto := <-canal //recebendo um valor
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto // mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)
}
