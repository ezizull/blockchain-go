package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func NewWallet() *Wallet {
	// 1. Creating ECDSA private key (32 bytes) public key (64 bytes)
	wallet := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	wallet.privateKey = privateKey
	wallet.publicKey = &wallet.privateKey.PublicKey

	// 2. Perform SHA-256 hashing on the public key (32 bytes)
	hash2 := sha256.New()
	hash2.Write(wallet.publicKey.X.Bytes())
	hash2.Write(wallet.publicKey.Y.Bytes())
	digest2 := hash2.Sum(nil)

	// 3. Perform RIPEMD-160 hashing on the result of SHA-256 (20 bytes)
	hash3 := ripemd160.New()
	hash3.Write(digest2)
	digest3 := hash3.Sum(nil)

	// 4. Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
	verByte4 := make([]byte, 21)
	verByte4[0] = 0x00
	copy(verByte4[1:], digest3[:])

	// 5. Perform SHA-256 hash on the extended RIPEMD-160 result
	hash5 := sha256.New()
	hash5.Write(verByte4)
	digest5 := hash5.Sum(nil)

	// 6. Perform SHA-256 hash on the result of the previous SHA-256 hash
	hash6 := sha256.New()
	hash6.Write(digest5)
	digest6 := hash6.Sum(nil)

	// 7. Take the first 4 bytes of the second SHA-256 hash. This is the address checksum
	chsum := digest6[:4]

	// 8. Add the 4 checksum bytes from stage 7 at the end of extended RIPEMD-160 hash from (25 bytes)
	checkByte8 := make([]byte, 25)
	copy(checkByte8[:21], verByte4[:])
	copy(checkByte8[21:], chsum[:])

	// 9. Convert the result from a byte string into a base58
	address := base58.Encode(checkByte8)
	wallet.blockChainAddress = address

	return wallet
}

func (wallet *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return wallet.privateKey
}

func (wallet *Wallet) PrivateKeyString() string {
	return fmt.Sprintf("%x", wallet.privateKey.D.Bytes())
}

func (wallet *Wallet) PublicKey() *ecdsa.PublicKey {
	return wallet.publicKey
}

func (wallet *Wallet) PublicKeyString() string {
	return fmt.Sprintf("%x%x", wallet.publicKey.X.Bytes(), wallet.privateKey.Y.Bytes())
}

func (wallet *Wallet) BlockChainAddress() string {
	return wallet.blockChainAddress
}
