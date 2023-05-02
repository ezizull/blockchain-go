package block

import (
	"fmt"
	"strings"
)

func (blockChain *BlockChain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	block := newBlock(nonce, prevHash, blockChain.transactionPool)
	blockChain.chain = append(blockChain.chain, block)
	blockChain.transactionPool = []*Transaction{}
	return block
}

func (blockChain *BlockChain) LastBlock() *Block {
	return blockChain.chain[len(blockChain.chain)-1]
}

func (blockChain *BlockChain) AddTranscation(sender string, recipient string, value float32) {
	transaction := NewTransaction(sender, recipient, value)
	blockChain.transactionPool = append(blockChain.transactionPool, transaction)
}

func (blockChain *BlockChain) CopyTransactionPool() []*Transaction {
	transctions := make([]*Transaction, 0)
	for _, trans := range blockChain.transactionPool {
		transctions = append(transctions,
			NewTransaction(trans.senderBlockChainAddress,
				trans.recipentBlockChainAddress,
				trans.value))
	}
	return transctions
}

func (blockChain *BlockChain) ValidProof(nonce int, prevHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, prevHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (blockChain *BlockChain) ProofOfWork() int {
	transactions := blockChain.CopyTransactionPool()
	previousHash := blockChain.LastBlock().Hash()
	nonce := 0
	for !blockChain.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func NewBlockChain() *BlockChain {
	block := &Block{}
	blockChain := new(BlockChain)
	blockChain.CreateBlock(0, block.Hash())
	return blockChain
}