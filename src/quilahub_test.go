package src

import (
	"io/ioutil"
	"testing"
)

func TestCreateHubDatabase(t *testing.T) {
	var createAquilaDb = DataStructCreateDb{
		Schema: SchemaStruct{
			Description: "this is my database",
			Unique:      "r8and0mseEd901",
			Encoder:     "strn:msmarco-distilbert-base-tas-b",
			Codelen:     768,
			Metadata: MetadataStructCreateDb{
				Name: "string",
				Age:  "number",
			},
		},
	}

	// wallet
	priv, err := ioutil.ReadFile("/home/dev/aquilax/ossl/private_unencrypted.pem")
	if err != nil {
		t.Error(err)
	}
	walletInitStruct := NewWallet(string(priv[:]))
	walletSign, err := walletInitStruct.CreateSignatureWallet(createAquilaDb)
	if err != nil {
		t.Error(err)
	}
	walletInitStruct.SecretKey = walletSign

	result, err := NewAquilaHub(walletInitStruct).CreateDatabase(&createAquilaDb, "http://localhost:5002/prepare")
	if err != nil {
		t.Error("Something went wrong.", err)
	}
	t.Log("=== Create Database Hub ==========================")
	t.Logf("%+v", result)
	t.Log("=== Create Database Hub ==========================")
}

// func TestCompressDocument(t *testing.T) {

// 	aquilaHubRequest := AquilaHubRequestStruct{
// 		Data: AquilaDataRequestStruct{
// 			Text:         []string{"It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.)"},
// 			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
// 		},
// 	}

// 	// wallet
// 	priv, err := ioutil.ReadFile("/home/dev/aquilax/ossl/private_unencrypted.pem")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	walletInitStruct := NewWallet(string(priv[:]))
// 	walletSign, err := walletInitStruct.CreateSignatureWallet(aquilaHubRequest)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	walletInitStruct.SecretKey = walletSign

// 	result, err := NewAquilaHub(walletInitStruct).CompressDocument(&aquilaHubRequest, "http://localhost:5002/compress")
// 	if err != nil {
// 		t.Error("Something went wrong.", err)
// 	}
// 	t.Log("=== Compress Document Hub ==========================")
// 	t.Logf("%+v", result)
// 	t.Log("=== Compress Document Hub ==========================")
// }
