package block

const MINING_DIFFICULTY = 3

type Block struct {
	timeStamp    int64
	nonce        int
	previousHash [32]byte
	transaction  []*Transaction
}
