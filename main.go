package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDosaque float64) string {
	podeSacar := valorDosaque > 0 && valorDosaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDosaque
		return "Saque realizado com sucesso"
	} else {
		return "Salado insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status, valor)
}

// Nessa aula:
// Criamos um nova conta corrente utilizando a palavra new;

// Em seguida, comparamos os tipos criados comparando suas referências e entendemos o que são ponteiros na prática;

// Para finalizar, desenvolvemos o método sacar que verifica se o valor do saque é maior do que zero e se a conta possui saldo.
