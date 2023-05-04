package main

import (
	"blockchain-go/domain/block"
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
	// blockChain.Mining()
	// blockChain.Print()

	// blockChain.AddTranscation("C", "D", 2.0)
	// blockChain.AddTranscation("X", "Y", 3.0)
	// blockChain.Mining()
	// blockChain.Print()

	// fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	// fmt.Printf("C  %.1f\n", blockChain.CalculateTotalAmount("C"))
	// fmt.Printf("D  %.1f\n", blockChain.CalculateTotalAmount("D"))

	// log.SetPrefix("Wallet: ")

	// wallt := wallet.NewWallet()
	// fmt.Println(wallt.PrivateKeyString())
	// fmt.Println(wallt.PublicKeyString())
	// fmt.Println(wallt.BlockChainAddress())

	// trans := wallet.NewTransaction(wallt.PrivateKey(), wallt.PublicKey(), wallt.BlockChainAddress(), "B", 1.0)
	// fmt.Printf("signature %s\n", trans.GenerateSignature())

	log.SetPrefix("Wallet Miner: ")

	walletMiner := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// wallet
	trans := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0)

	// blockchain
	blockChain := block.NewBlockChain(walletMiner.BlockChainAddress())
	isAdded := blockChain.AddTranscation(walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0, walletA.PublicKey(), trans.GenerateSignature())
	fmt.Println("Added ", isAdded)

	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("A %.1f\n", blockChain.CalculateTotalAmount(walletA.BlockChainAddress()))
	fmt.Printf("B %.1f\n", blockChain.CalculateTotalAmount(walletB.BlockChainAddress()))
	fmt.Printf("Miner %.1f\n", blockChain.CalculateTotalAmount(walletMiner.BlockChainAddress()))
}
