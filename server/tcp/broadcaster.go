package tcp

import (
	"encoding/json"
	"github.com/rossus/simple-blockchain/blockchain"
	"io"
	"log"
	"net"
	"time"
)

func broadcast(conn net.Conn) {
	for {

		time.Sleep(30 * time.Second)

		output, err := json.Marshal(blockchain.GetValue())
		if err != nil {
			log.Println(">>>>---->")
			log.Fatal(err)
		}

		_, err = io.WriteString(conn, string(output))
		if err != nil {
			log.Println(">>||---->")
			log.Fatal(err)
		}

	}
}
