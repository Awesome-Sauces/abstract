package state

import "github.com/Awesome-Sauces/abstract/abstract/blockchain"

/*

This is meant for interfacing with the StateMachine between the App
which then will interface with the p2p layer

*/

func (st *StateRuntime) AddTransaction(tx *blockchain.Transaction) {
	st.currentBlock.AddTransaction(tx)
}
