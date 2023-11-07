package state

import (
	"log"
	"testing"

	"github.com/Awesome-Sauces/abstract/abstract/blockchain"
	"github.com/Awesome-Sauces/abstract/crypto"
)

func TestState(t *testing.T) {
	t.Run("PaymentTest", func(t *testing.T) {

		privatekey, err := crypto.DecodePrivateKey("0x689086063b6edcc373ff61e689f8cc88f12d14ef73feefe5a54d6db70bc840e0")
		kp := crypto.NewKeyPair()

		if err != nil {
			log.Println(err)
		}

		state := Start("nan")

		state.Genesis()

		block := blockchain.NewBlock("", 0)

		block.Issuer = crypto.NewKeyPair().Address

		tx := *blockchain.NewPublicLedgerFundTransfer(privatekey.PublicKey().Address(), kp.Address, 30.25)

		tx.Hash = tx.GenerateHash()

		tx_sig, err := privatekey.EncodeSignature(tx.Hash)

		if err != nil {
			log.Fatal(err)
		}

		tx.Signature = tx_sig

		block.AddTransaction(tx)

		block.Hash = block.GenerateHash()

		sig, err := privatekey.EncodeSignature("block.Hash")

		if err != nil {
			log.Fatal(err)
		}

		block.AddSignature(privatekey.PublicKey().Address(), sig)

		err = state.TransactionSettlement(block)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("TRANSACTION SUCCESS")

		log.Println(state.GetACE(kp.Address))
		log.Println(state.GetACE(privatekey.PublicKey().Address()))

		state.PrintAllACE()

	})
}
