package src

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateHubDatabase(t *testing.T) {

	err := LoadEnvFile()
	if err != nil {
		t.Fatal("Fail to load .env file. ", err)
	}

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

	walletInitStruct, err := CreateWalletSignForTesting(createAquilaDb)
	if err != nil {
		t.Error("Something went wrong.", err)
	}

	url := fmt.Sprintf("http://%v:%v/prepare",
		os.Getenv("AQUILA_DB_HOST"),
		os.Getenv("AQUILA_HUB_PORT"),
	)

	result, err := NewAquilaHub(walletInitStruct).CreateDatabase(&createAquilaDb, url)
	if err != nil {
		t.Error("Something went wrong.", err)
	}
	t.Log("=== Create Database Hub ==========================")
	t.Logf("%+v", result)
	t.Log("=== Create Database Hub ==========================")
}

func TestCompressDocument(t *testing.T) {

	err := LoadEnvFile()
	if err != nil {
		t.Fatal("Fail to load .env file. ", err)
	}

	aquilaHubRequest := AquilaHubRequestStruct{
		Data: AquilaDataRequestStruct{
			Text:         []string{"It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.)"},
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
	}

	walletInitStruct, err := CreateWalletSignForTesting(aquilaHubRequest)
	if err != nil {
		t.Error("Something went wrong.", err)
	}

	url := fmt.Sprintf("http://%v:%v/compress",
		os.Getenv("AQUILA_DB_HOST"),
		os.Getenv("AQUILA_HUB_PORT"),
	)

	result, err := NewAquilaHub(walletInitStruct).CompressDocument(&aquilaHubRequest, url)
	if err != nil {
		t.Error("Something went wrong.", err)
	}
	t.Log("=== Compress Document Hub ==========================")
	t.Logf("%+v", result)
	t.Log("=== Compress Document Hub ==========================")
}
