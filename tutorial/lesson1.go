package tutorial

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/rossus/simple_blockchain/blockchain"
	s_http "github.com/rossus/simple_blockchain/server/http"
	"github.com/rossus/simple_blockchain/types"
	"log"
	"time"
)

func lesson1main() {
	log.Println("Starting simple blockchain lesson 1 app")

	err := godotenv.Load()
	if err != nil {
		log.Println("ioioo---->")
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := types.Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		blockchain.Append(genesisBlock)
	}()

	log.Fatal(s_http.Run())
}
