package src

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/mr-tron/base58/base58"
	"gopkg.in/mgo.v2/bson"
)

type WalletStruct struct {
	SecretKey string
}

func Wallet() WalletStruct {
	return WalletStruct{}
}

// create signature
func CreateSignWallet(requestStructure interface{}) string {

	bson_data, err := bson.Marshal(requestStructure)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%x\n", bson_data) // Right !!!

	bytes := sha512.Sum384(bson_data)
	// fmt.Printf("%x\n", bytes) // Right !!!

	priv, _ := ioutil.ReadFile("/home/dev/aquilax/ossl/private_unencrypted.pem") // ???
	privPem, _ := pem.Decode(priv)
	privPemBytes := privPem.Bytes

	priv1, _ := x509.ParsePKCS1PrivateKey([]byte(privPemBytes))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, priv1, crypto.SHA384, bytes[:]) // Right !!!

	signedHash := base58.Encode(signature)
	// fmt.Println("==========================")
	// fmt.Println(num) // Right !!!!!!

	return signedHash
}
