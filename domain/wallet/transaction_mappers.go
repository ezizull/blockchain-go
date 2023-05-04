package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"

	signDomain "blockchain-go/domain/signature"
)

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string,
	recipient string, value float32) *Transaction {
	return &Transaction{privateKey, publicKey, sender, recipient, value}
}

func (trans *Transaction) GenerateSignature() *signDomain.Signature {
	marshal, _ := json.Marshal(trans)
	hash := sha256.Sum256([]byte(marshal))
	r, s, _ := ecdsa.Sign(rand.Reader, trans.senderPrivateKey, hash[:])
	return &signDomain.Signature{r, s}
}

func (trans *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		value     float32 `json:"value`
	}{
		Sender:    trans.senderBlockChainAddress,
		Recipient: trans.recipientBlockChainAddress,
		value:     trans.value,
	})
}
