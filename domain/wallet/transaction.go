package wallet

import "crypto/ecdsa"

type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockChainAddress    string
	recipientBlockChainAddress string
	value                      float32
}
