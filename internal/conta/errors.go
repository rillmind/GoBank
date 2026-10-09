package conta

import "errors"

var (
	ErrSaldoInsuficiente  = errors.New("Saldo insuficiente para realizar operação.")
	ErrValorInvalido      = errors.New("Valor inválido.")
	ErrContaNaoEncontrada = errors.New("Conta não encontrada.")
	ErrTimeOut            = errors.New("Operação inválida devido a tempo limite. Aguarde.")
)
