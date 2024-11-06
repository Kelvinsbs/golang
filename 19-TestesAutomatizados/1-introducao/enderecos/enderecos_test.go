package enderecos_test //unica forma de ter mais de um pacote

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // nao e obrigatorio
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Rodovia joão gualberto soares", "Rodovia"},
		{"Avenida josé alberto", "Avenida"},
		{"Praça da Sé", "Tipo Inválido"},
		{"Estrada olivinhas", "Estrada"},
		{"RUA DOS MALUCOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel() // nao e obrigatorio

	if 1 > 2 {
		t.Errorf("Teste Quebrou!")
	}
}

// TESTE MAIS SIMPLES POSSIVEL
// func TestTipoDeEndereco(t *testing.T) {
// 	enderecoParaTeste := "Rua ABC"

// 	tipoDeEnderecoEsperado := "Avenida"

// 	tipoDeEnderecoRecedido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecedido != tipoDeEnderecoEsperado {
// 		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
// 			tipoDeEnderecoEsperado, tipoDeEnderecoRecedido)
// 	}
// }
