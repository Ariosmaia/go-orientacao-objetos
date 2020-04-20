package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// * Para apontar para conta de quem estiver chamando
func (c *ContaCorrente) Sacar(valorDosaque float64) string {
	podeSacar := valorDosaque > 0 && valorDosaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDosaque
		return "Saque realizado com sucesso"
	} else {
		return "Salado insuficiente"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println(contaDaSilvia.saldo)
}
