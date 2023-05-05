package block

import (
	"encoding/hex"
	"fmt"
)

// ToBlockResponseArray
func ToBlockResponseArray(blocks []*Block) []*BlockResponse {
	responses := make([]*BlockResponse, len(blocks))
	for i, block := range blocks {
		previousHashStr := hex.EncodeToString(block.previousHash[:])
		response := &BlockResponse{
			TimeStamp:    block.timeStamp,
			Nonce:        block.nonce,
			PreviousHash: previousHashStr,
			Transaction:  block.transaction,
		}
		responses[i] = response
	}
	return responses
}

// Print
func (block *Block) Print() {
	fmt.Printf("timestamp       %d\n", block.timeStamp)
	fmt.Printf("nonce           %d\n", block.nonce)
	fmt.Printf("previous_hash   %x\n", block.previousHash)
	for _, trans := range block.transaction {
		trans.Print()
	}
}
