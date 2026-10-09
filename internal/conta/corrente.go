package conta

import (
	"fmt"

	"github.com/rillmind/GoBank/internal/cpf"
)

type Corrente struct {
	Nome                 string
	Cpf                  cpf.CPF
	Numero               string
	Saldo                int
	SaldoLimite          int
	SaldoChequeEspecial  int
	LimiteChequeEspecial int
	Tarifa               float64
	Tipo                 Tipo
}

func (cc *Corrente) GetUserInfo() string {
	return fmt.Sprintf("Numero: %s, Titular: %s.", cc.Numero, cc.Nome)
}

func (cc *Corrente) GetSaldo() float64 {
	return float64(cc.Saldo) / 100
}

func (cc *Corrente) GetTipo() Tipo {
	return cc.Tipo
}

func (cc *Corrente) Depositar(valor int) error {
	cc.Saldo += valor
	return nil
}

func (cc *Corrente) Sacar(valor int) error {
	cc.Saldo -= valor
	return nil
}
