package interfaces

type IBankAccount interface {
	GetBalance() int
	DepositBalance(amount int) (int, error)
	WithdrawBalance(amount int) (int, error)
}
