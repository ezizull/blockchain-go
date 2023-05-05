package security

import (
	"io/ioutil"

	"github.com/golang-jwt/jwt/v5"
)

const (
	privatePath = "application/security/keys/app.rsa"
	publicPath  = "application/security/keys/app.rsa.pub"
)

func GettingKeySSH() (err error) {
	privateBytes, err := ioutil.ReadFile(privatePath)
	if err != nil {
		return
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return
	}

	publicBytes, err := ioutil.ReadFile(publicPath)
	if err != nil {
		return
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return
	}

	return
}
