package bchainer

import (
	"github.com/rossus/simple_blockchain/types"
	"time"
)

func GenerateBlock(oldBlock types.Block, BPM int) (types.Block, error) {

	var newBlock types.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index+1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}
