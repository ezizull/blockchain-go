package block

import (
	"fmt"
	"strings"
)

func (block *Block) Print() {
	fmt.Printf("timestamp       %d\n", block.timeStamp)
	fmt.Printf("none            %d\n", block.nonce)
	fmt.Printf("previous_hash   %s\n", block.previousHash)
	fmt.Printf("transaction     %s\n", block.transaction)
}

func (blockChain *BlockChain) Print() {
	fmt.Println()

	for i, block := range blockChain.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("-", 15), i, strings.Repeat("-", 15))
		block.Print()
	}

	fmt.Println()
}
