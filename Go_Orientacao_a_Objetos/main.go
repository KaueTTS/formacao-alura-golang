package main

import (
	"fmt"

	"Go_Orientacao_a_Objetos/clientes"
	"Go_Orientacao_a_Objetos/contas"
)

func PagarBoleto(conta verificaConta, valor float64) string {
	return conta.Sacar(valor)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteJoao := clientes.Titular{
		Nome:      "João Silva",
		Sexo:      "Masculino",
		CPF:       "123.456.789-00",
		Profissao: "Engenheiro",
		Idade:     30,
	}

	contaJoao := contas.ContaCorrente{
		Titular:       clienteJoao,
		NumeroConta:   12345,
		NumeroAgencia: 678,
		Ativo:         true,
	}

	contaJoao.Depositar(1000.00)
	fmt.Println("Saldo inicial:", contaJoao.ObterSaldo())

	fmt.Println(PagarBoleto(&contaJoao, 150.00))
}
