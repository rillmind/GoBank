# Changelog - GoBank

Todas as mudanças notáveis deste projeto serão documentadas neste arquivo.

## [Não commitado] - 2026-10-09

### Adicionado
- `internal/conta/poupanca.go`: novo tipo `Poupanca` (`Nome, Numero, Cpf, Saldo uint, Tipo, TimeOut time.Duration`)
  com construtor `NewPoupanca()` e métodos `GetUserInfo(), GetSaldo(), GetTipo(), Depositar(), Sacar(), RenderJuros(taxa float64)`.
  - `Depositar` valida `valor < 0` e retorna `ErrValorInvalido` (bloco de `TimeOut` deixado como TODO comentado).
  - `RenderJuros` aplica `Saldo += Saldo * taxa`.
- `internal/conta/errors.go`: erros centralizados `ErrSaldoInsuficiente, ErrValorInvalido, ErrContaNaoEncontrada, ErrTimeOut`.
- `internal/banco/iConta.go`: nova interface `Rendivel { RenderJuros(taxa float64) }`.
- `internal/conta/corrente.go`: novo construtor `NewCorrente(nome, numero, cpf, saldo, limiteChequeEspecial, tarifa, tipo)`.

### Alterado
- `internal/conta/corrente.go`:
  - `struct Corrente`: removidos `SaldoLimite` e `SaldoChequeEspecial`; ordem ajustada para `Nome, Numero, Cpf, Saldo, LimiteChequeEspecial, Tarifa, Tipo`.
  - `Depositar(valor int)` -> `Depositar(valor float64)`: converte reais para centavos com `int(valor * 100)`.
  - `Sacar(valor int)` -> `Sacar(valor float64)`: desconta `valor + Tarifa`, faz rollback e retorna `ErrSaldoInsuficiente` se `Saldo <= LimiteChequeEspecial`.

### Observações
- `banco.Conta` está desatualizada (ainda declara `GetSaldo() int`, `Depositr(valor int)`, `Sacar(valor int)`, `GetContaTipo()`) e não é satisfeita por `Corrente/Poupanca` (`GetSaldo() float64`, `Depositar/Sacar(float64)`, `GetTipo()`).
- `Rendivel.RenderJuros` não retorna `error`, mas `Poupanca.RenderJuros` retorna `error`.
- `Poupanca.Sacar` verifica `if Saldo < 0` sobre `uint` (condição inalcançável) e faz underflow antes de validar.

## [f438b08] - 2026-10-09 - Inittial Commit: esqueleto incompleto. conta corrente pre pronta.

### Adicionado
- `cmd/banco/main.go`: `func main()` vazio.
- `go.mod`: módulo `github.com/rillmind/GoBank`, go `1.26.8`.
- `internal/banco/banco.go`: pacote `banco` vazio.
- `internal/banco/iConta.go`: interface `Conta { GetUserInfo() string; GetSaldo() int; GetContaTipo() conta.Tipo; Depositr(valor int) error; Sacar(valor int) error }`.
- `internal/conta/tipo.go`: tipo `Tipo int` com `TipoCorrente, TipoPoupanca`.
- `internal/conta/corrente.go`: `struct Corrente` inicial (`Nome, Cpf, Numero, Saldo, SaldoLimite, SaldoChequeEspecial, LimiteChequeEspecial, Tarifa, Tipo`) com `GetUserInfo(), GetSaldo() float64, GetTipo(), Depositar(valor int), Sacar(valor int)` sem validações.
- `internal/cpf/cpf.go`: `struct CPF { numero string }` com `Validar()` e `Formatar()` como TODO.
