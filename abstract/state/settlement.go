package state

import (
	"log"

	"github.com/Awesome-Sauces/abstract/abstract/blockchain"
	"github.com/Awesome-Sauces/abstract/crypto"
	"github.com/Awesome-Sauces/abstract/ozone"
)

/*

The settlement and validation functions
do not care about the state of the blockchain
instead the commit and register functions will worry about such,
by implementing a complex system for tracking multiple blocks by
using a registrar that contains blocks that are ready for consensus
but have not yet been distributed. This system may not be needed
but will definitely be a nice backup plan

*/

// Logic for validating a block, this means the transactions
// contained and the signatures on the block as-well
func (st *StateRuntime) validateBlock(block blockchain.Block) bool {
	for _, tx := range block.Transactions {
		if !st.validateTransaction(*tx) {
			return false
		}
	}

	if !block.VerifyHash() {
		return false
	}

	for address, signature := range block.Signatures {
		pk, err := crypto.RecoverPublicKeyFromSignature(block.Hash, signature)

		if err != nil {
			return false
		}

		if address != pk.Address() {
			return false
		}
	}

	return true
}

// Logic for Verifiying the Validity of a Transaction
func (st *StateRuntime) validateTransaction(tx blockchain.Transaction) bool {

	if !tx.VerifyHash() {
		return false
	}

	pk, kerr := crypto.RecoverPublicKeyFromSignature(tx.Hash, tx.Signature)

	if kerr != nil {
		return false
	}

	if pk.Address() != tx.From {
		return false
	}

	ledger, err := st.getLedger(tx.Domain)

	gas_fee := 0.0001

	amount, _ := tx.Amount.Float64()

	if tx.AuthToken != "NAN" {
		gas_fee = 0.001
	}

	if tx.Domain != "abstract.pub" {
		gas_fee += 0.0001
	}

	if err != nil {
		return false
	}

	if tx.Domain == "abstract.pub" && st.aceBalance(tx.From)-(gas_fee+amount) > 0.0 {
		return true
	} else if st.aceBalance(tx.From)-gas_fee >= 0.0 && ledger.GetBalance(tx.From)-amount >= 0.0 {
		return true
	}

	return false
}

func (st *StateRuntime) settleBlock(block blockchain.Block) (bool, string) {
	if !st.validateBlock(block) {
		return false, "<BLOCK-FAILED-VALIDATION>"
	}

	for _, tx := range block.Transactions {
		result, rout := st.settleTransaction(block.Issuer, *tx)

		if !result {
			return result, rout
		}
	}

	return true, "<BLOCK-SETTLEMENT-SUCCESS><NO-FAILED-TRANSACTIONS>"
}

// Logic for Settling a single-transaction
// Assumes the logic for transaction
func (st *StateRuntime) settleTransaction(administrator string, tx blockchain.Transaction) (bool, string) {
	// Return output false = fail true = success, string = metadata
	ledger, err := st.getLedger(tx.Domain)

	gas_fee := 0.0001

	amount, _ := tx.Amount.Float64()

	if tx.AuthToken != "NAN" {
		gas_fee = 0.001
	}

	if tx.Domain != "abstract.pub" {
		gas_fee += 0.0001
	}

	if err != nil {
		return false, "<LEDGER-REQUEST-FAILED>"
	}

	if tx.Domain == "abstract.pub" && st.aceBalance(tx.From)-(gas_fee+amount) > 0.0 {
		result := ledger.FundsTransfer(tx.From, tx.To, amount, 100)

		core, err := st.getLedger("abstract.pub")

		if err != nil {
			return false, "<TRANSACTION-FAILURE><UNKOWN-OPERATION>"
		}

		core.FundsTransfer(tx.From, administrator, gas_fee, 100)

		if result != nil {
			return false, "<TRANSACTION-FAILURE><UNKOWN-OPERATION>"
		}

		return true, "<TRANSACTION-SUCCESS><FUNDS-TRANSFERRED>"
	} else if st.aceBalance(tx.From)-gas_fee >= 0.0 && ledger.GetBalance(tx.From)-amount >= 0.0 {
		result := ledger.FundsTransfer(tx.From, tx.To, amount, 100)

		core, err := st.getLedger("abstract.pub")

		if err != nil {
			return false, "<TRANSACTION-FAILURE><UNKOWN-OPERATION>"
		}

		core.FundsTransfer(tx.From, administrator, gas_fee, 100)

		if result != nil {
			return false, "<TRANSACTION-FAILURE><UNKOWN-OPERATION>"
		}

		return true, "<TRANSACTION-SUCCESS><FUNDS-TRANSFERRED>"
	}

	return false, "<TRANSACTION-FAILURE><INSUFFICIENT-BALANCE>"
}

func (st *StateRuntime) ValidateChain(chain string) bool {
	side_chain := ozone.New(1000)

	err := side_chain.LoadString(chain)

	if err != nil {
		log.Print(err)

		return false
	}

	return true
}
