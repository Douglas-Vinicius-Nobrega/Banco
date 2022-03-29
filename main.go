package main

import (
	"fmt"

	"github.com/Alura/Banco/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}

	contaExemplo.Depositar(1001)

	fmt.Println(contaExemplo.ObterSaldo())

}
