package blockchain

import (
	"encoding/json"
	"math/big"
)

// Account struct
type Account struct {
	Address string     `json:"address"`
	Balance *big.Float `json:"balance"`
}

// NewAccount creates a new Account instance.
func NewAccount(address string, balance float64) *Account {
	return &Account{
		Address: address,
		Balance: big.NewFloat(balance),
	}
}

func (acc *Account) Set(amount float64) {
	acc.Balance = big.NewFloat(amount)
}

// Credit adds the given amount to the account balance.
func (acc *Account) Credit(amount float64) {
	if amount > 0 {
		acc.Balance.Add(acc.Balance, big.NewFloat(amount))
	}
}

// Debit subtracts the given amount from the account balance and returns true if successful.
func (acc *Account) Debit(amount float64) bool {
	if amount > 0 {
		subAmount := big.NewFloat(amount)
		if acc.Balance.Cmp(subAmount) >= 0 {
			acc.Balance.Sub(acc.Balance, subAmount)
			return true
		}
	}
	return false
}

// FromBytes deserializes an Account from a byte slice.
func AccountFromBytes(data []byte) (*Account, error) {
	var acc Account
	if err := json.Unmarshal(data, &acc); err != nil {
		return nil, err
	}
	return &acc, nil
}

// ToBytes serializes an Account to a byte slice.
func AccountToBytes(acc *Account) ([]byte, error) {
	data, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
