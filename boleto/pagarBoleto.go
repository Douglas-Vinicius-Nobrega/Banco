package boleto

func PagarBoleto(conta VerificarConta, valorDoBoleto float64) { // conta : primeiro parametro, do tipo interface que vai chamar, VerificarConta, valorBoleto do tipo float64
	conta.Sacar(valorDoBoleto)

}

type VerificarConta interface {
	Sacar(valor float64) string
}
