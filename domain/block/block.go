package block

type BlockChain struct {
	transPool []string
	chain     []*Block
}

type Block struct {
	nonce        int
	previousHash string
	timeStamp    int64
	transaction  []string
}
