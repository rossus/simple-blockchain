package tcp

import "github.com/rossus/simple-blockchain/types"

var bcServer chan []types.Block

func SetChannel(newch chan []types.Block) {
	bcServer = newch
}

func GetChannel() chan []types.Block {
	return bcServer
}