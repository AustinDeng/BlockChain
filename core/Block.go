package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// 区块结构体定义
type Block struct {
	Index int64
	TimeStamp int64
	PrevBlockHash string
	Hash string
	Data string
}

// 生成新的区块
func CreateNewBlock(pre Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = pre.Index + 1
	newBlock.TimeStamp = time.Now().Unix()
	newBlock.PrevBlockHash = pre.Hash
	newBlock.Hash = calculateHash(newBlock)
	newBlock.Data = data
	return newBlock
}

// 计算区块的 Hash 值
func calculateHash(block Block) string {
	data := string(block.Index) + string(block.TimeStamp) + block.PrevBlockHash + block.Data
	blockInBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(blockInBytes[:])
}

// 生成创世区块
func CreateGenesisBlock() Block{
	block := Block{}
	block.Index = -1
	block.Hash = ""
	return CreateNewBlock(block, "Genesis Block")
}