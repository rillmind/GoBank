package conta

import (
	"fmt"
	"time"

	"github.com/rillmind/GoBank/internal/cpf"
)

type Poupanca struct {
	Nome    string
	Numero  string
	Cpf     cpf.CPF
	Saldo   uint
	Tipo    Tipo
	TimeOut time.Duration
}

func NewPoupanca(
	nome, numero string,
	cpf cpf.CPF,
	saldo uint,
	tipo Tipo,
	timeOut time.Duration,
) *Poupanca {
	return &Poupanca{
		nome,
		numero,
		cpf,
		saldo,
		tipo,
		timeOut,
	}
}

func (cp *Poupanca) GetUserInfo() string {
	return fmt.Sprintf("Numero: %s, Titular: %s.", cp.Numero, cp.Nome)
}

func (cp *Poupanca) GetSaldo() float64 {
	return float64(cp.Saldo) / 100
}

func (cp *Poupanca) GetTipo() Tipo {
	return cp.Tipo
}

func (cp *Poupanca) Depositar(valor float64) error {
	if valor < 0 {
		return ErrValorInvalido
	}

	// TODO: sistema de timeout
	// 	if cp.TimeOut != 0 {
	// 		return ErrTimeOut
	// 	}
	//
	// 	cp.TimeOut += 15 * time.Second

	cp.Saldo += uint(valor * 100)

	return nil
}

func (cp *Poupanca) Sacar(valor float64) error {
	cp.Saldo -= uint(valor * 100)

	if cp.Saldo < 0 {
		cp.Saldo += uint(valor * 100)
		return ErrSaldoInsuficiente
	}

	return nil
}

func (cp *Poupanca) RenderJuros(taxa float64) error {
	cp.Saldo += uint(float64(cp.Saldo) * taxa)
	return nil
}
