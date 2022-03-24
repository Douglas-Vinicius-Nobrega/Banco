package main

import "fmt"

type contaCorrente struct { // tipo de estrutura... uma estrutura que tem vários campos dentro dele
	 titular string 
	 numeroAgencia int 
	 numeroConta int 
	 saldo float64
}

func (c * contaCorrente) Sacar(valorDoSaque float64) string { // quem solicitar a func Sacar, eu vou olhar para a conta que está chamando (c * contaCorrente), nesse caso, não precisamos especificar a conta de quem ira sacar
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo // valor do saque tem que ser maior que 0 E o valor do saque tem que ser menor ou igual o valor do saldo
	if podeSacar { // se for verdade, que tem saldo
		c.saldo -= valorDoSaque // subtraindo o valor que tem no saque 
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insulficiente para saque"
	}
}

func main() {
	contaDaSilvia := contaCorrente {
		titular: "Silva",
		saldo: 500.00,
	}

	fmt.Println("Saldo:", contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println("Saldo restante:", contaDaSilvia.saldo)

}