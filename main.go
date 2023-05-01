package main

import (
	"blockchain-go/domain/block"
	"log"
)

func main() {
	log.SetPrefix("Blockchain: ")

	blockChain := block.NewBlockChain()
	blockChain.Print()

	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()

	blockChain.CreateBlock(4, "hash 2")
	blockChain.Print()
}
