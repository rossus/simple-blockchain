package blockchain

import "github.com/rossus/simple_blockchain/types"

var Blockchain []types.Block

func GetValue() []types.Block {
	return Blockchain
}

func GetItself() *[]types.Block {
	return &Blockchain
}

func Append(block types.Block) {
	Blockchain = append(Blockchain, block)
}