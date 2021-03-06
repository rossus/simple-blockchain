package tutorial

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/rossus/simple-blockchain/blockchain"
	s_tcp "github.com/rossus/simple-blockchain/server/tcp"
	"github.com/rossus/simple-blockchain/types"
	"log"
	"time"
)

func lesson2main() {
	log.Println("Starting simple blockchain lesson 2 app")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	s_tcp.SetChannel(make(chan []types.Block))

	t := time.Now()
	genesisBlock := types.Block{0, t.String(), 0, "", ""}
	spew.Dump(genesisBlock)
	blockchain.Append(genesisBlock)

	s_tcp.Run()
}
