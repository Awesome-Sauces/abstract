package state

import (
	"errors"
	"log"
	"math/big"
	"strconv"

	"github.com/Awesome-Sauces/abstract/abstract/blockchain"
	"github.com/Awesome-Sauces/abstract/ozone"
)

// StateRuntime manages the state of the blockchain by handling blocks and transactions.
type StateRuntime struct {
	currentBlock *blockchain.Block

	ledgers    *ozone.Database
	blockchain *ozone.Database
	index      *big.Int
}

// Start creates a new instance of StateRuntime and initializes the required databases.
func Start(filename string) *StateRuntime {
	return &StateRuntime{
		ledgers:    ozone.New(1000),
		blockchain: ozone.New(1000),
		index:      big.NewInt(0),
	}
}

// **DISCLAIMER - THIS FUNCTION WILL HEAVILY EAT AWAY RESOURCES
// **USE ONLY IN DIRE CIRCUMSTANCE (CHARGE HIGH GAS FEE TO USE, e.g. 10000)

// FetchTransaction retrieves a transaction from the blockchain based on its hash.
func (st *StateRuntime) FetchTransaction(hash string) (*blockchain.Transaction, error) {
	var tx *blockchain.Transaction

	err := st.blockchain.Iterate(func(key string, value ozone.DataItem) error {
		block := value.Value.(*blockchain.Block)

		if t, contains := block.Transactions[hash]; contains {
			tx = t
			return nil
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tx, nil
}

// FetchBlock retrieves a block from the blockchain based on its index.
func (st *StateRuntime) FetchBlock(index int) (*blockchain.Block, error) {
	block, err := st.blockchain.Get(strconv.Itoa(index))
	if err != nil {
		return nil, err
	}

	return block.Value.(*blockchain.Block), nil
}

// Genesis initializes the blockchain with a genesis block and a default ledger.
func (st *StateRuntime) Genesis() error {
	abstractLedger := blockchain.NewLedger()

	err := abstractLedger.RegisterAccount(blockchain.NewAccount("0xbb2e1e379f1cdaaf308f877fcc1b43d63d276fa9", 3000))
	if err != nil {
		log.Fatal(err)
	}

	err = st.registerLedger("abstract.pub", abstractLedger)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (st *StateRuntime) getLedger(domain string) (*blockchain.Ledger, error) {
	dti, err := st.ledgers.Get(domain)

	return dti.Value.(*blockchain.Ledger), err
}

func (st *StateRuntime) aceBalance(address string) float64 {
	ledger, err := st.getLedger("abstract.pub")

	if err != nil {
		return 0.0
	}

	return ledger.GetBalance(address)
}

func (st *StateRuntime) registerLedger(domain string, ledger *blockchain.Ledger) error {
	if _, err := st.ledgers.Get(domain); err != nil {
		err := st.ledgers.Set(domain, ledger)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("ledger already exists under domain")
}

// GetBalance returns the balance of an account in a specific ledger.
func (st *StateRuntime) GetBalance(ledgerName string, address string) (float64, error) {
	ledger, err := st.ledgers.Get(ledgerName)
	if err != nil {
		return 0.0, err
	}

	return ledger.Value.(*blockchain.Ledger).GetBalance(address), nil
}
