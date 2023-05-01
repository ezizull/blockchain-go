package block

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (trans *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("sender_blockchain_address      %s\n", trans.senderBlockChainAddress)
	fmt.Printf("repicent_blockchain_address    %s\n", trans.recipentBlockChainAddress)
	fmt.Printf("value                          %.1f\n", trans.value)
}

func (trans *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_marshal"`
		Recipient string  `json:"recipient_blockchain_marshal"`
		Value     float32 `json:"value"`
	}{
		Sender:    trans.senderBlockChainAddress,
		Recipient: trans.recipentBlockChainAddress,
		Value:     trans.value,
	})
}
