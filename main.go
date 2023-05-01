package main

import (
	"blockchain-go/domain/block"
	"log"
)

func main() {
	log.SetPrefix("Blockchain: ")

	blockChain := block.NewBlockChain()
	blockChain.Print()

	blockChain.AddTranscation("A", "B", 1.0)
	prevHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, prevHash)
	blockChain.Print()

	blockChain.AddTranscation("C", "D", 2.0)
	blockChain.AddTranscation("X", "Y", 3.0)
	prevHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(4, prevHash)
	blockChain.Print()
}
