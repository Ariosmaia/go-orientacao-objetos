package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

// Uma interface é a definição de um conjunto de métodos comuns a um ou mais tipos. É o que permite a criação de tipos polimórficos em Go.
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDennis := contas.ContaPoupanca{}
	contaDoDennis.Depositar(100)
	PagarBoleto(&contaDoDennis, 60)

	fmt.Println(contaDoDennis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())
}

// Nessa aula:
// Criamos um novo tipo de conta: a ContaPoupança;

// Para finalizar, criamos um novo tipo interface onde podemos utilizar tanto a conta corrente como poupança para pagar um boleto através da função Sacar.
