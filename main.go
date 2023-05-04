package main

import (
	"blockchain-go/domain/wallet"
	"fmt"
	"log"
)

func main() {
	// log.SetPrefix("Blockchain: ")

	// myBlockChainAddress := "my_blockchain_address"

	// blockChain := block.NewBlockChain(myBlockChainAddress)
	// blockChain.Print()

	// blockChain.AddTranscation("A", "B", 1.0)
	// blockChain.Maining()
	// blockChain.Print()

	// blockChain.AddTranscation("C", "D", 2.0)
	// blockChain.AddTranscation("X", "Y", 3.0)
	// blockChain.Maining()
	// blockChain.Print()

	// fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	// fmt.Printf("C  %.1f\n", blockChain.CalculateTotalAmount("C"))
	// fmt.Printf("D  %.1f\n", blockChain.CalculateTotalAmount("D"))

	log.SetPrefix("Wallet: ")

	wallet := wallet.NewWallet()
	fmt.Println(wallet.PrivateKeyString())
	fmt.Println(wallet.PublicKeyString())
	fmt.Println(wallet.BlockChainAddress())
}
