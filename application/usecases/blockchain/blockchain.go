package blockchain

import (
	blockDomain "blockchain-go/domain/block"
	walletDomain "blockchain-go/domain/wallet"
)

// Service is a struct that contains the repository implementation for block use case
type Service struct {
}

// GetAllChainByPort is a function that returns blockchain by port
func (s *Service) GetAllChainByPort(port uint64) (*blockDomain.BlockchainResponse, error) {
	walletMiner := walletDomain.NewWallet()
	blockChain := blockDomain.NewBlockChain(walletMiner.BlockChainAddress(), port)
	return blockChain.ToBlockChainResponse(), nil
}
