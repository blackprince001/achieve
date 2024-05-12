package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	UserID  uint
	User    User     `gorm:"foreignkey:UserID"`
	Wallets []Wallet `json:"wallets"`
}

type WalletType int

const (
	Goal WalletType = iota + 1
	Savings
)

type Wallet struct {
	gorm.Model
	AccountID      uint
	Account        Account `gorm:"foreignkey:AccountID"`
	Name           string
	WalletType     WalletType
	CurrentBalance float64
	Target         float64
	IsLocked       bool
}

type TransactionType int

const (
	Deposit TransactionType = iota + 1
	Withdrawal
)

type TransactionStatus string

const (
	Pending   TransactionStatus = "pending"
	Completed TransactionStatus = "completed"
	Failed    TransactionStatus = "failed"
)

type Transaction struct {
	gorm.Model
	WalletID  uint
	Wallet    Wallet `gorm:"foreignkey:WalletID"`
	Type      TransactionType
	Amount    float64
	Status    TransactionStatus
	StatusMsg string
}
