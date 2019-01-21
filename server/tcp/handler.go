package tcp

import (
	"bufio"
	"github.com/davecgh/go-spew/spew"
	"github.com/rossus/simple_blockchain/bchainer"
	"github.com/rossus/simple_blockchain/blockchain"
	"github.com/rossus/simple_blockchain/types"
	"io"
	"log"
	"net"
	"strconv"
)

var bcServer = make(chan []types.Block)

func handleConn(conn net.Conn) {

	defer conn.Close()

	_, err := io.WriteString(conn, "Enter a new BPM:")
	if err != nil {
		log.Println("/:---->")
		log.Fatal(err)
	}

	Blockchain := blockchain.GetValue()

	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			bpm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Println("-=-=------>")
				log.Printf("%v is not a number: %v", scanner.Text(), err)
				continue
			}

			newBlock, err := bchainer.GenerateBlock(Blockchain[len(Blockchain)-1], bpm)
			if err != nil {
				log.Println("yuyuyu---->")
				log.Println(err)
				continue
			}

			if bchainer.IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				bchainer.ReplaceChain(newBlockchain)
			}

			bcServer <- Blockchain

			_, err = io.WriteString(conn, "\nEnter a new BPM:\n")
			if err != nil {
				log.Println("32rewr---->")
				log.Fatal(err)
			}
		}
	}()

	go broadcast(conn)

	for _ = range bcServer {
		spew.Dump(blockchain.GetValue())
	}
}
