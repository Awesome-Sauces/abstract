package blockchain

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Block struct {
	PreviousHash string
	Transactions map[string]*Transaction // Change the value type to *Transaction
	Hash         string
	TxAmount     int32
	GasAmount    *big.Float // Change the type to *big.Float
	Signatures   map[string]string
	Index        int
	Issuer       string
}

func NewBlock(previousHash string, index int) *Block {
	return &Block{
		PreviousHash: previousHash,
		Transactions: make(map[string]*Transaction), // Use *Transaction
		Index:        index,
		Signatures:   make(map[string]string),
		GasAmount:    new(big.Float).SetFloat64(0.0), // Initialize GasAmount with a big.Float
	}
}

func (block *Block) AddTransaction(tx *Transaction) { // Change the parameter type to *Transaction
	if _, contains := block.Transactions[tx.Hash]; !contains {
		// Convert GasAllocated from big.Int to big.Float before adding
		gasFloat := new(big.Float).SetInt(tx.GasAllocated)
		block.GasAmount.Add(block.GasAmount, gasFloat) // Use big.Float arithmetic
		block.Transactions[tx.Hash] = tx
		block.TxAmount++
	}
}

func (block *Block) AddSignature(address string, signature string) {
	if _, contains := block.Signatures[address]; !contains {
		block.Signatures[address] = signature
	}
}

func (block *Block) VerifyHash() bool { // Change to pointer receiver
	return block.Hash == block.GenerateHash()
}

func (block *Block) SetHash() { // Change to pointer receiver
	block.Hash = block.GenerateHash()
}

func (block *Block) GenerateHash() string { // Change to pointer receiver
	// Concatenating the PreviousHash, Signatures, and Index
	dataToHash := block.PreviousHash

	for _, transaction := range block.Transactions {
		dataToHash += transaction.Signature
	}

	dataToHash += fmt.Sprint(block.Index)

	// Creating a hash of the concatenated string
	hash := sha512.New()
	hash.Write([]byte(dataToHash))
	hashInBytes := hash.Sum(nil)

	// Converting the hash to a hexadecimal string
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}
