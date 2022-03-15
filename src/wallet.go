package src

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"

	"github.com/mr-tron/base58/base58"
	"gopkg.in/mgo.v2/bson"
)

type WalletStruct struct {
	SecretKey string
}

func NewWallet(privKey string) WalletStruct {
	return WalletStruct{
		SecretKey: privKey,
	}
}

// create signature
func (w *WalletStruct) CreateSignatureWallet(requestStructure interface{}) (string, error) {

	bson_data, err := bson.Marshal(requestStructure)
	if err != nil {
		return "", err
	}
	// fmt.Println("+++++++++++++++++++++++++++++++++++")
	// fmt.Printf("%x\n", bson_data) // Right !!!

	bytes := sha512.Sum384(bson_data)
	// fmt.Printf("%x\n", bytes) // Right !!!

	// priv, err := ioutil.ReadFile("/home/dev/aquilax/ossl/private_unencrypted.pem") // ???
	// if err != nil {
	// 	return "", err
	// }
	privPem, _ := pem.Decode([]byte(w.SecretKey))
	privPemBytes := privPem.Bytes

	priv1, err := x509.ParsePKCS1PrivateKey([]byte(privPemBytes))
	if err != nil {
		return "", err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv1, crypto.SHA384, bytes[:]) // Right !!!
	if err != nil {
		return "", err
	}

	signedHash := base58.Encode(signature)
	// fmt.Println("====================================")
	// fmt.Println(signedHash) // Right !!!!!!

	return signedHash, nil
}
