package main

import (
	"io"
	"encoding/json"
	"net/http"
	"BlockChain/core"
)

// 定义一个全局的区块链对象
var blockchain *core.BlockChain

func main() {
	blockchain = core.CreateNewBlockChain()
	run()
}

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandle)
	http.HandleFunc("/blockchain/write", blockchainWriteHandle)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// 获取区块链数据
func blockchainGetHandle(w http.ResponseWriter, r *http.Request){
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// 写入区块链数据
func blockchainWriteHandle(w http.ResponseWriter, r *http.Request){
	blockData := r.URL.Query().Get("data")
	blockchain.SetData(blockData)
	blockchainGetHandle(w, r)
}
