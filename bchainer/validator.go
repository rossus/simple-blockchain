package bchainer

import (
	"github.com/rossus/simple-blockchain/types"
)

func IsBlockValid(newBlock, oldBlock types.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
