package block

type Block struct {
	timeStamp    int64
	transaction  []*Transaction
	nonce        int
	previousHash [32]byte
}
