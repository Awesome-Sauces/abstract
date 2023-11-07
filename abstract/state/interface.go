package state

import (
	"log"
	"math/big"

	"github.com/Awesome-Sauces/abstract/abstract/blockchain"
)

/*

This is meant for interfacing with the StateMachine between the App
which then will interface with the p2p layer

*/

func (st *StateRuntime) Balance(domain string, address string) *big.Float {
	ledger, err := st.getLedger(domain)

	if err != nil {
		return big.NewFloat(0.0)
	}

	return big.NewFloat(ledger.GetBalance(address))
}

func (st *StateRuntime) AddTransaction(tx *blockchain.Transaction) {
	st.currentBlock.AddTransaction(tx)
}

func (st *StateRuntime) CommitBlock(block *blockchain.Block) {
	if approved := st.validateBlock(*block); approved {
		if approved, result := st.settleBlock(*block); approved {
			log.Println(result)
		}
	}
}

func (st *StateRuntime) GetBlockchain() string {
	blockchain, err := st.blockchain.SaveString()

	if err != nil {
		log.Print(err)
	}

	return blockchain
}
