package core

import (
	"log"
	"sync"
	"time"

	"github.com/perlin-network/noise"
)

var Blockchain []Block
var Node *noise.Node
var DataChannel chan Block
var Mutex *sync.Mutex

func init() {
	Mutex = new(sync.Mutex)
	DataChannel = make(chan Block)

	Blockchain = make([]Block, 0)
	genesisBlock := Block{}
	genesisBlock = Block{
		Index:       0,
		Timestamp:   time.Now().String(),
		SavingMoney: 0,
		Hash:        genesisBlock.Hashing(),
		PrevHash:    "",
	}
	Blockchain = append(Blockchain, genesisBlock)

	log.Println(Blockchain)
}
