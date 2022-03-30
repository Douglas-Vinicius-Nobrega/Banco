package main

import (
	"fmt"
	"time"

	"github.com/Alura/Banco/boleto"
	"github.com/Alura/Banco/contas"
)

func main() {

	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(500)

	boleto.PagarBoleto(&contaDenis, 50)
	contaLuisa := contas.ContaCorrente{}
	contaLuisa.Depositar(1000)

	boleto.PagarBoleto(&contaLuisa, 600)

	fmt.Println(contaDenis.ObterSaldo(), contaLuisa.ObterSaldo())

	data := time.Now()
	fmt.Println(data)
}
