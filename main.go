package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/rossus/simple_blockchain/blockchain"
	"github.com/rossus/simple_blockchain/server"
	"github.com/rossus/simple_blockchain/types"
	"log"
	"time"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := types.Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		blockchain.Append(genesisBlock)
	}()

	log.Fatal(server.Run())
}
