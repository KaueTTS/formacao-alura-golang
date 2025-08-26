package contas

import (
	"fmt"

	"Go_Orientacao_a_Objetos/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroConta, NumeroAgencia int
	Ativo                      bool
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Novo saldo: " + fmt.Sprintf("%.2f", c.saldo)
	} else {
		return "Saque não realizado. Verifique o valor e o saldo."
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso! Novo saldo: " + fmt.Sprintf("%.2f", c.saldo)
	} else {
		return "Depósito não realizado. O valor deve ser positivo."
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.saldo += valorDaTransferencia
		return "Transferência realizada com sucesso! Novo saldo: " + fmt.Sprintf("%.2f", c.saldo)
	} else {
		return "Transferência não realizada. Verifique o valor e o saldo."
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
