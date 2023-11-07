package blockchain

import (
	"errors"
	"log"

	"github.com/Awesome-Sauces/abstract/ozone"
)

var (
	errEntryExists         = errors.New("account already registered")
	errFundsTransferFailed = errors.New("funds transfer failed to execute, no changes made")
	errInsufficientFunds   = errors.New("not enough funds present to complete transfer")
)

type Ledger struct {
	index    int
	accounts *ozone.Database
}

func NewLedger() *Ledger {
	return &Ledger{
		index:    0,
		accounts: ozone.New(100000), // Specifying cache size as 1000, adjust as needed
	}
}

/*

CORE FUNCTIONS

*/

/*

UTILITY FUNCTIONS

*/

func (ledger *Ledger) FetchBalances() map[string]Account {
	rout := make(map[string]Account)

	err := ledger.accounts.Iterate(func(key string, value ozone.DataItem) error {

		account, err := AccountFromBytes(value.Value.([]byte))
		if err != nil {
			return err
		}

		rout[key] = *account

		return nil
	})

	if err != nil {
		log.Println(err)
		return rout
	}

	return rout
}

func (ledger *Ledger) RegisterAccount(account *Account) error {
	_, err := ledger.accounts.Get(account.Address)

	if err == ozone.ErrKeyNotFound {
		value, err := AccountToBytes(account)
		if err != nil {
			return err
		}

		return ledger.accounts.Set(account.Address, value)
	} else if err == nil {
		return errEntryExists
	}

	return err
}

func (ledger *Ledger) FundsTransfer(from string, to string, amount float64, offset int) error {
	_, errFrom := ledger.accounts.Get(from)
	_, errTo := ledger.accounts.Get(to)

	if errFrom != nil {
		return errFundsTransferFailed
	}

	balance := ledger.GetBalance(from)
	if balance-amount >= 0 {
		if errTo != nil {
			ledger.RegisterAccount(NewAccount(to, 0.0))
		}

		err := ledger.SetBalance(from, balance-amount)
		if err != nil {
			return err
		}

		err = ledger.SetBalance(to, ledger.GetBalance(to)+(amount*(float64(offset)/100)))
		if err != nil {
			_ = ledger.SetBalance(from, balance)
			return errFundsTransferFailed
		}

		return nil
	}

	return errFundsTransferFailed
}

func (ledger *Ledger) SetBalance(address string, amount float64) error {
	dataItem, err := ledger.accounts.Get(address)

	if err != nil {
		return err
	}

	account, err := AccountFromBytes(dataItem.Value.([]byte))
	if err != nil {
		return err
	}

	account.Set(amount)

	bytesVal, err := AccountToBytes(account)
	if err != nil {
		return err
	}

	return ledger.accounts.Set(address, bytesVal)
}

func (ledger Ledger) GetBalance(address string) float64 {
	dataItem, err := ledger.accounts.Get(address)
	if err != nil {
		return 0.0
	}

	account, err := AccountFromBytes(dataItem.Value.([]byte))
	if err != nil {
		return 0.0
	}

	val, _ := account.Balance.Float64()

	return val
}

// Other methods like LedgerToBytes and LedgerFromBytes remain the same as they don't interact directly with the ozone DB.
