package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDaSilvia.Transferir(-200, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}

// Nessa aula:
// Criamos um método chamado depositar, que quando invocado nos retorna uma mensagem (string) e o valor do novo salário (float64);

// Em seguida, criamos o método transferir, que tira um valor de uma conta e transfere para uma conta destino reutilizando o método depositar;

// Para finalizar, criamos um pacote chamado contas e criamos um arquivo chamado contaCorrente.go para manter todo código referente a este tipo de conta.
