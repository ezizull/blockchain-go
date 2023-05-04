package blockchain

import (
	blockDomain "blockchain-go/domain/block"
	walletDomain "blockchain-go/domain/wallet"
	"fmt"
	"log"
)

func blockChain() {
	log.SetPrefix("Blockchain: ")

	// wallet
	walletA := walletDomain.NewWallet()
	walletB := walletDomain.NewWallet()
	walletX := walletDomain.NewWallet()

	myBlockChainAddress := "my_blockchain_address"
	blockChain := blockDomain.NewBlockChain(myBlockChainAddress)
	blockChain.Print()

	// transfer
	transA := walletDomain.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0)

	blockChain.AddTranscation("A", "B", 1.0, walletA.PublicKey(), transA.GenerateSignature())
	blockChain.Mining()
	blockChain.Print()

	// transfer
	transB := walletDomain.NewTransaction(walletB.PrivateKey(), walletB.PublicKey(), walletB.BlockChainAddress(), walletX.BlockChainAddress(), 2.0)
	transX := walletDomain.NewTransaction(walletX.PrivateKey(), walletX.PublicKey(), walletX.BlockChainAddress(), walletA.BlockChainAddress(), 3.0)

	blockChain.AddTranscation("B", "X", 2.0, walletB.PublicKey(), transB.GenerateSignature())
	blockChain.AddTranscation("X", "A", 3.0, walletX.PublicKey(), transX.GenerateSignature())
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C  %.1f\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("D  %.1f\n", blockChain.CalculateTotalAmount("D"))
}

func wallet() {
	log.SetPrefix("Wallet: ")

	wallt := walletDomain.NewWallet()
	fmt.Println(wallt.PrivateKeyString())
	fmt.Println(wallt.PublicKeyString())
	fmt.Println(wallt.BlockChainAddress())

	trans := walletDomain.NewTransaction(wallt.PrivateKey(), wallt.PublicKey(), wallt.BlockChainAddress(), "B", 1.0)
	fmt.Printf("signature %s\n", trans.GenerateSignature())

}

func walletMiner() {
	log.SetPrefix("Wallet Miner: ")

	walletMiner := walletDomain.NewWallet()
	walletA := walletDomain.NewWallet()
	walletB := walletDomain.NewWallet()

	// wallet
	trans := walletDomain.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0)

	// blockchain
	blockChain := blockDomain.NewBlockChain(walletMiner.BlockChainAddress())
	isAdded := blockChain.AddTranscation(walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0, walletA.PublicKey(), trans.GenerateSignature())
	fmt.Println("Added ", isAdded)

	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("A %.1f\n", blockChain.CalculateTotalAmount(walletA.BlockChainAddress()))
	fmt.Printf("B %.1f\n", blockChain.CalculateTotalAmount(walletB.BlockChainAddress()))
	fmt.Printf("Miner %.1f\n", blockChain.CalculateTotalAmount(walletMiner.BlockChainAddress()))
}
