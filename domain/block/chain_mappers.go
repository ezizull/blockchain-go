package block

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

func NewBlockChain() *BlockChain {
	block := &Block{}
	blockChain := new(BlockChain)
	blockChain.CreateBlock(0, block.Hash())
	return blockChain
}
