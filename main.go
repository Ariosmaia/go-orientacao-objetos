package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoDennis := contas.ContaPoupanca{}
	contaDoDennis.Depositar(100)

	fmt.Println(contaDoDennis.ObterSaldo())
}
