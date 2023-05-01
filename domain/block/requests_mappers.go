package block

import "time"

func newBlock(nonce int, prevHash string) *Block {
	block := new(Block)
	block.timeStamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = prevHash
	return block
}

func (blockChain *BlockChain) CreateBlock(nonce int, prevHash string) *Block {
	block := newBlock(nonce, prevHash)
	blockChain.chain = append(blockChain.chain, block)
	return block
}

func NewBlockChain() *BlockChain {
	blockChain := new(BlockChain)
	blockChain.CreateBlock(0, "init hash")
	return blockChain
}
