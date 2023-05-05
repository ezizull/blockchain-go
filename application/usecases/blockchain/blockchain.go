package blockchain

import (
	blockDomain "blockchain-go/domain/block"
	walletDomain "blockchain-go/domain/wallet"
)

// Service is a struct that contains the repository implementation for block use case
type Service struct {
}

// GetByPort is a function that returns blockchain by port
func (s *Service) GetByPort(port uint64) (*blockDomain.BlockChain, error) {
	walletMiner := walletDomain.NewWallet()
	blockChain := blockDomain.NewBlockChain(walletMiner.BlockChainAddress(), port)
	return blockChain, nil
}
