package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// newBlock
func newBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	block := new(Block)
	block.timeStamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = prevHash
	block.transaction = transactions
	return block
}

// Hash
func (block *Block) Hash() [32]byte {
	marshal, _ := json.Marshal(block)
	return sha256.Sum256([]byte(marshal))
}

// MarshalJSON
func (block *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash string         `json:"previous_hash"`
		Transaction  []*Transaction `json:"transaction"`
	}{
		Timestamp:    block.timeStamp,
		Nonce:        block.nonce,
		PreviousHash: fmt.Sprintf("%x", block.previousHash),
		Transaction:  block.transaction,
	})
}

// ToBlockResponse
func (block *Block) ToBlockResponse() *BlockResponse {
	return &BlockResponse{
		TimeStamp:    block.timeStamp,
		Nonce:        block.nonce,
		PreviousHash: fmt.Sprintf("%x", block.previousHash),
		Transaction:  block.transaction,
	}
}
