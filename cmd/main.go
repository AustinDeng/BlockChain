package main

import "BlockChain/core"

func main() {
	bc := core.CreateNewBlockChain()
	bc.SetData("This is the first block.")
	bc.SetData("This is the second block.")
	bc.Print()
}