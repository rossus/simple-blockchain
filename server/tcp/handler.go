package tcp

import (
	"bufio"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/rossus/simple-blockchain/bchainer"
	"github.com/rossus/simple-blockchain/blockchain"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func handleConn(conn net.Conn) {

	defer conn.Close()

	_, err := io.WriteString(conn, "Enter a new BPM:")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			Blockchain := blockchain.GetValue()

			bpm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v is not a number: %v", scanner.Text(), err)
				continue
			}

			newBlock, err := bchainer.GenerateBlock(Blockchain[len(Blockchain)-1], bpm)
			if err != nil {
				log.Println(err)
				continue
			}

			if bchainer.IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				bchainer.ReplaceChain(newBlockchain)
			}

			bcServer <- blockchain.GetValue()

			_, err = io.WriteString(conn, "\nEnter a new BPM:")
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(30 * time.Second)

			output, err := json.Marshal(blockchain.GetValue())
			if err != nil {
				log.Fatal(err)
			}

			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(blockchain.GetValue())
	}
}
