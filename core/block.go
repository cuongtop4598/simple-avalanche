package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index       int    `json:"index"`
	Timestamp   string `json:"timestamp"`
	SavingMoney int    `json:"saving_money"`
	Hash        string `json:"hash"`
	PrevHash    string `json:"prev_hash"`
}

func NewBlock(prevBlock *Block, savingMoney int) *Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.SavingMoney = savingMoney
	newBlock.PrevHash = prevBlock.Hash

	newBlock.Hash = newBlock.Hashing()

	return &newBlock
}

func (b *Block) Hashing() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.SavingMoney) + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) IsValid(prevBlock *Block) bool {
	if prevBlock.Index+1 != b.Index {
		return false
	}

	if prevBlock.Hash != b.PrevHash {
		return false
	}

	if b.Hashing() != b.Hash {
		return false
	}

	return true
}
