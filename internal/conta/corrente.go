package conta

import (
	"fmt"

	"github.com/rillmind/GoBank/internal/cpf"
)

type Corrente struct {
	Nome                 string
	Numero               string
	Cpf                  cpf.CPF
	Saldo                int
	LimiteChequeEspecial int
	Tarifa               float64
	Tipo                 Tipo
}

func NewCorrente(
	nome, numero string,
	cpf cpf.CPF,
	saldo, limiteChequeEspecial int,
	tarifa float64,
	tipo Tipo,
) *Corrente {
	return &Corrente{
		nome,
		numero,
		cpf,
		saldo,
		limiteChequeEspecial,
		tarifa,
		tipo,
	}
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

func (cc *Corrente) Depositar(valor float64) error {
	cc.Saldo += int(valor * 100)
	return nil
}

func (cc *Corrente) Sacar(valor float64) error {
	cc.Saldo -= int(valor*100) + int(cc.Tarifa*100)

	if cc.Saldo <= cc.LimiteChequeEspecial {
		cc.Saldo += int(valor*100) + int(cc.Tarifa*100)
		return ErrSaldoInsuficiente
	}

	return nil
}
