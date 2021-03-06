package http

import (
	"encoding/json"
	"github.com/rossus/simple-blockchain/bchainer"
	"github.com/rossus/simple-blockchain/blockchain"
	"github.com/rossus/simple-blockchain/types"
	"io"
	"log"
	"net/http"
)

func HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(blockchain.GetValue(), "", "  ")

	if err != nil {
		log.Println("##---->")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.WriteString(w, string(bytes))
	if err != nil {
		log.Println("///---->")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m types.Message
	Blockchain := blockchain.GetValue()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := bchainer.GenerateBlock(Blockchain[len(Blockchain)-1], m.BPM)
	if err != nil {
		log.Println("<><>---->")
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	if bchainer.IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		newBlockchain := append(Blockchain, newBlock)
		bchainer.ReplaceChain(newBlockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}