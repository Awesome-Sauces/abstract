package blockchain

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Transaction struct {
	From         string     `json:"from"`
	To           string     `json:"to"`
	Amount       *big.Float `json:"amount"`
	Type         string     `json:"type"`
	Domain       string     `json:"domain"`
	AuthToken    string     `json:"authtoken"`
	Codec        string     `json:"codec"`
	Architecture string     `json:"architecture"`
	GasAllocated *big.Int   `json:"gas"`
	Hash         string     `json:"hash"`
	Signature    string     `json:"signature"`
}

func (transaction Transaction) VerifyHash() bool {
	return transaction.Hash == transaction.GenerateHash()
}

func (transaction *Transaction) SetHash() {
	transaction.Hash = transaction.GenerateHash()
}

func (transaction Transaction) GenerateHash() string {
	// Concatenating the fields to be hashed
	dataToHash := fmt.Sprintf("%s%s%s%s%s%s%s%s%s",
		transaction.From,
		transaction.To,
		transaction.Amount.String(),
		transaction.Type,
		transaction.Domain,
		transaction.AuthToken,
		transaction.Codec,
		transaction.Architecture,
		transaction.GasAllocated.String(),
	)

	// Creating a hash of the concatenated string
	hash := sha512.New()
	hash.Write([]byte(dataToHash))
	hashInBytes := hash.Sum(nil)

	// Converting the hash to a hexadecimal string
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}

/*

FUND_TRANSFER

REQUIRES : DOMAIN NAME WITH A REGISTERED LEDGER

*/

func NewPrivateLedgerFundTransfer(from string, to string, amount *big.Float, domain string, authToken string) *Transaction {
	return &Transaction{
		From:         from,
		To:           to,
		Amount:       amount,
		Type:         "FUND_TRANSFER",
		Domain:       domain,
		AuthToken:    authToken,
		Codec:        "NAN",
		Architecture: "NAN",
		GasAllocated: big.NewInt(10000),
	}
}

func NewPublicLedgerFundTransfer(from string, to string, amount *big.Float) *Transaction {
	return &Transaction{
		From:         from,
		To:           to,
		Amount:       amount,
		Type:         "FUND_TRANSFER",
		Domain:       "abstract.pub",
		AuthToken:    "NAN",
		Codec:        "NAN",
		Architecture: "NAN",
		GasAllocated: big.NewInt(10),
	}
}

func NewCustomLedgerFundTransfer(from string, to string, amount *big.Float, domain string) *Transaction {
	return &Transaction{
		From:         from,
		To:           to,
		Amount:       amount,
		Type:         "FUND_TRANSFER",
		Domain:       domain,
		AuthToken:    "NAN",
		Codec:        "NAN",
		Architecture: "NAN",
		GasAllocated: big.NewInt(10),
	}
}

func NewProgramCreate(from string, architecture string, codec string) *Transaction {
	return &Transaction{
		From:         from,
		To:           "AVM",
		Amount:       big.NewFloat(0.0),
		Type:         "PROGRAM_CREATE",
		Domain:       "abstract.pub",
		AuthToken:    "NAN",
		Codec:        codec,
		Architecture: architecture,
		GasAllocated: big.NewInt(10),
	}
}
