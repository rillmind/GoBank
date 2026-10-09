package banco

import (
	"github.com/rillmind/GoBank/internal/conta"
)

type Conta interface {
	GetUserInfo() string
	GetSaldo() int
	GetContaTipo() conta.Tipo
	Depositr(valor int) error
	Sacar(valor int) error
}
