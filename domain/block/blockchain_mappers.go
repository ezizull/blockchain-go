package block

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	signDomain "blockchain-go/domain/signature"
)

func (blockChain *BlockChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"chains"`
	}{
		Blocks: blockChain.chain,
	})
}

func (blockChain *BlockChain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	block := newBlock(nonce, prevHash, blockChain.transactionPool)
	blockChain.chain = append(blockChain.chain, block)
	blockChain.transactionPool = []*Transaction{}
	return block
}

func (blockChain *BlockChain) LastBlock() *Block {
	return blockChain.chain[len(blockChain.chain)-1]
}

func (blockChain *BlockChain) AddTranscation(sender string, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, sign *signDomain.Signature) bool {
	transaction := NewTransaction(sender, recipient, value)

	if sender == MINING_SENDER {
		blockChain.transactionPool = append(blockChain.transactionPool, transaction)
		return true
	}

	if blockChain.VerifyTransactionSignature(senderPublicKey, sign, transaction) {
		// if blockChain.CalculateTotalAmount(sender) < value {
		// 	log.Println("ERROR: No enough balance in a wallet")
		// 	return false
		// }

		blockChain.transactionPool = append(blockChain.transactionPool, transaction)
		return true
	} else {
		log.Println("ERROR: Verify Transaction")
		return false
	}
}

func (blockChain *BlockChain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, sign *signDomain.Signature,
	trans *Transaction) bool {
	marshal, _ := json.Marshal(trans)
	hash := sha256.Sum256([]byte(marshal))
	return ecdsa.Verify(senderPublicKey, hash[:], sign.R, sign.S)
}

func (blockChain *BlockChain) CopyTransactionPool() []*Transaction {
	transctions := make([]*Transaction, 0)
	for _, trans := range blockChain.transactionPool {
		transctions = append(transctions,
			NewTransaction(trans.senderBlockChainAddress,
				trans.recipientBlockChainAddress,
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

func (blockChain *BlockChain) Mining() bool {
	blockChain.AddTranscation(MINING_SENDER, blockChain.address, MINING_REWARD, nil, nil)
	nonce := blockChain.ProofOfWork()
	prevHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(nonce, prevHash)
	log.Printf("action=mining, status=success\n\n")
	return true
}

func (blockChain *BlockChain) CalculateTotalAmount(blockChainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, block := range blockChain.chain {
		for _, trans := range block.transaction {
			value := trans.value
			if blockChainAddress == trans.recipientBlockChainAddress {
				totalAmount += value
			}

			if blockChainAddress == trans.senderBlockChainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

func NewBlockChain(blockChainAddress string, port uint64) *BlockChain {
	block := new(Block)
	blockChain := new(BlockChain)
	blockChain.address = blockChainAddress
	blockChain.CreateBlock(0, block.Hash())
	blockChain.port = port
	return blockChain
}
