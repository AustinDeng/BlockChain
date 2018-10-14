package core

import(
	"fmt"
	"log"
)

// 区块链结构体定义
type BlockChain struct {
	Block []*Block
}

// 生成区块链
func CreateNewBlockChain() *BlockChain {
	genesisBlock := CreateGenesisBlock()
	blockChain := BlockChain{}
	blockChain.AppendBlock(&genesisBlock)
	return &blockChain
}

// 给 BlockChain 添加 Block
func (bc *BlockChain) AppendBlock(b *Block){
	if len(bc.Block) == 0 {
		bc.Block = append(bc.Block, b)
		return
	}
	if isValid(*b, *bc.Block[len(bc.Block) - 1]) {
		bc.Block = append(bc.Block, b)
		return
	} else {
		log.Fatal("Invalid Block.")
	}

}

// 给 BlockChain 添加给定值的 Block
func (bc *BlockChain) SetData(data string){
	preBlock := bc.Block[len(bc.Block) - 1]
	currentBlock := CreateNewBlock(*preBlock, data)
	bc.AppendBlock(&currentBlock)
}

// 打印 Block Chain 数据
func (bc *BlockChain) Print() {
	for _,block := range bc.Block {
		fmt.Printf("Index: %d \n",block.Index)
		fmt.Printf("PrevBlockHash: %s \n",block.PrevBlockHash)
		fmt.Printf("CurrHash: %s \n",block.Hash)
		fmt.Printf("Data: %s \n",block.Data)
		fmt.Printf("Timestamp: %d \n",block.TimeStamp)
		fmt.Println()
	}
}

// 校验传入的 Block 是否有效
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	return true
}
