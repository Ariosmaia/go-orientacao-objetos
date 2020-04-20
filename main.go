package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.123.123-12", "Desenvolvedor GO"}
	contadoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	fmt.Println(contadoBruno)
}
