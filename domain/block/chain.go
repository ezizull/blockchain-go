package block

type BlockChain struct {
	transactionPool []*Transaction
	chain           []*Block
	address         string
}
