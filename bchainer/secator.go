package bchainer

import (
	"github.com/rossus/simple_blockchain/blockchain"
	"github.com/rossus/simple_blockchain/types"
)

func ReplaceChain(newBlocks []types.Block) {
	bchain := blockchain.GetItself()
	if len(newBlocks) > len(*bchain) {
		*bchain = newBlocks
	}
}
