package contas

import (
	"Go_Orientacao_a_Objetos/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroConta, NumeroAgencia, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Novo saldo: " + fmt.Sprintf("%.2f", c.saldo)
	} else {
		return "Saque não realizado. Verifique o valor e o saldo."
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso! Novo saldo: " + fmt.Sprintf("%.2f", c.saldo)
	} else {
		return "Depósito não realizado. O valor deve ser positivo."
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
