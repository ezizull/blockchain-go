package block

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Block struct {
	timeStamp    int64
	nonce        int
	previousHash [32]byte
	transaction  []*Transaction
}

type BlockResponse struct {
	TimeStamp    int64          `json:"timestamp"`
	Nonce        int            `json:"nonce"`
	PreviousHash string       `json:"previous_hash"`
	Transaction  []*Transaction `json:"transaction"`
}
