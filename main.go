package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}

// Nessa aula:
// Criamos um novo pacote chamado clientes e um arquivo chamado clientes.go, onde desenvolvemos uma nova struct chamada Titular com nome, CPF e profissão;

// Em seguida, alteramos o campo titular da struct contaCorrente para ser do tipo da struct Titular;

// Para finalizar, alteramos a visibilidade do campo saldo e criamos um método chamado obterSaldo.
