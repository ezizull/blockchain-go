package block

import (
	"fmt"
	"strings"
)

// ToBlockChainResponse
func (blockChain *BlockChain) ToBlockChainResponse() *BlockchainResponse {
	return &BlockchainResponse{
		Chain: ToBlockResponseArray(blockChain.chain),
	}
}

// Print
func (blockChain *BlockChain) Print() {
	fmt.Println()
	for i, block := range blockChain.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("-", 15), i, strings.Repeat("-", 15))
		block.Print()
	}
	fmt.Println()
}
