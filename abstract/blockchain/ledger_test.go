package blockchain

import (
	"log"
	"testing"
)

func TestLedgerOperation(t *testing.T) {
	t.Run("PaymentTest", func(t *testing.T) {
		// Create a Ledger to store value
		ledger := NewLedger()

		// Create the Accounts we want
		account_1 := NewAccount("0x1Ae31370249Fa6DA54c60268bfCe8e0755E08E29", 0.0)
		account_2 := NewAccount("0xccd9A97D248a6D73a1AcB1dA28E9133939F1aC4f", 0.0)

		// Register 0x1Ae31370249Fa6DA54c60268bfCe8e0755E08E29 to the Ledger (Allows for Fund Transfer)
		err := ledger.RegisterAccount(account_1)
		if err != nil {
			log.Fatal(err)
		}

		// Register 0xccd9A97D248a6D73a1AcB1dA28E9133939F1aC4f to the Ledger (Allows for Fund Transfer)
		err = ledger.RegisterAccount(account_2)
		if err != nil {
			log.Fatal(err)
		}

		// Give 0x1Ae31370249Fa6DA54c60268bfCe8e0755E08E29 the necessary funds
		err = ledger.SetBalance("0x1Ae31370249Fa6DA54c60268bfCe8e0755E08E29", 32.13)

		if err != nil {
			log.Fatal(err)
		}

		// Initiate the Transfer with a value of 10
		err = ledger.FundsTransfer("0x1Ae31370249Fa6DA54c60268bfCe8e0755E08E29", "0xccd9A97D248a6D73a1AcB1dA28E9133939F1aC4f", 10.0, 100)

		if err != nil {
			log.Fatal(err)
		}

		// Verify Post-Payment Balances
		balance := ledger.GetBalance("0x1Ae31370249Fa6DA54c60268bfCe8e0755E08E29")

		log.Println("0x1Ae31370249Fa6DA54c60268bfCe8e0755E08E29", balance)

		balance = ledger.GetBalance("0xccd9A97D248a6D73a1AcB1dA28E9133939F1aC4f")

		log.Println("0xccd9A97D248a6D73a1AcB1dA28E9133939F1aC4f", balance)

	})
}
