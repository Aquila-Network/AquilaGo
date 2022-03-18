package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func CreateWalletSignForTesting(requestStructure interface{}) (WalletStruct, error) {
	var wallet WalletStruct
	pathUnencryptedPemFile := os.Getenv("PATH_TO_PRIVATE_UNENCRYPTED_PEM_FILE")
	priv, err := ioutil.ReadFile(pathUnencryptedPemFile)
	if err != nil {
		return wallet, err
	}
	walletInitStruct := NewWallet(string(priv[:]))
	walletSign, err := walletInitStruct.CreateSignatureWallet(requestStructure)
	if err != nil {
		return wallet, err
	}
	walletInitStruct.SecretKey = walletSign

	return walletInitStruct, nil
}

func TestCreateDatabase(t *testing.T) {

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

	url := fmt.Sprintf("http://%v:%v/db/create",
		os.Getenv("AQUILA_DB_HOST"),
		os.Getenv("AQUILA_DB_PORT"),
	)

	result, err := NewAquilaDb(walletInitStruct).CreateDatabase(&createAquilaDb, url)
	if err != nil {
		t.Error("Something went wrong.", err)
	}
	t.Log("=== Create database ==========================")
	t.Logf("%+v", result)
	t.Log("=== Create database ==========================")
}

func TestInsertDocument(t *testing.T) {

	err := LoadEnvFile()
	if err != nil {
		t.Fatal("Fail to load .env file. ", err)
	}

	var docInsert = DatatDocInsertStruct{
		Docs: []DocsStruct{
			{
				Payload: PayloadStruct{
					Metadata: MetadataStructDocInsert{
						Name: "name1",
						Age:  20,
					},
					Code: []float64{0.1, 0.2, 0.3},
				},
			},
			{
				Payload: PayloadStruct{
					Metadata: MetadataStructDocInsert{
						Name: "name1",
						Age:  20,
					},
					Code: []float64{0.1, 0.2, 0.3},
				},
			},
		},
		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
	}

	// req, err := json.Marshal(docInsert)
	// fmt.Println("=========================")
	// fmt.Print(string(req[:]))

	walletInitStruct, err := CreateWalletSignForTesting(docInsert)
	if err != nil {
		t.Error("Something went wrong.", err)
	}

	url := fmt.Sprintf("http://%v:%v/db/doc/insert",
		os.Getenv("AQUILA_DB_HOST"),
		os.Getenv("AQUILA_DB_PORT"),
	)

	result, err := NewAquilaDb(walletInitStruct).InsertDocument(&docInsert, url)
	if err != nil {
		t.Error("Something went wrong.", err.Error())
	}
	t.Log("=== Doc Insert ==========================")
	t.Logf("%+v", result)
	t.Log("=== Doc Insert ==========================")
}

func TestSearch(t *testing.T) {

	err := LoadEnvFile()
	if err != nil {
		t.Fatal("Fail to load .env file. ", err)
	}

	matrix := make([][]float64, 1)
	matrix[0] = make([]float64, 1)
	matrix[0] = []float64{
		-0.01806008443236351, -0.17380790412425995, 0.03992759436368942, 0.43514639139175415,
	}
	searchBody := DataSearchStruct{
		Matrix:       matrix,
		K:            10,
		R:            0,
		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
	}

	walletInitStruct, err := CreateWalletSignForTesting(searchBody)
	if err != nil {
		t.Error("Something went wrong.", err)
	}

	url := fmt.Sprintf("http://%v:%v/db/search",
		os.Getenv("AQUILA_DB_HOST"),
		os.Getenv("AQUILA_DB_PORT"),
	)

	result, err := NewAquilaDb(walletInitStruct).SearchKDocument(&searchBody, url)
	if err != nil {
		t.Error("Something went wrong.", err)
	}
	t.Log("=== Doc Search ==========================")
	t.Logf("%+v", result)
	t.Log("=== Doc Search ==========================")
}

func TestDeleteDocument(t *testing.T) {

	err := LoadEnvFile()
	if err != nil {
		t.Fatal("Fail to load .env file. ", err)
	}

	var docDelete = DeleteDataStruct{
		Ids: []string{
			"3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
			"BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ",
		},
		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
	}

	walletInitStruct, err := CreateWalletSignForTesting(docDelete)
	if err != nil {
		t.Error("Something went wrong.", err)
	}

	url := fmt.Sprintf("http://%v:%v/db/doc/delete",
		os.Getenv("AQUILA_DB_HOST"),
		os.Getenv("AQUILA_DB_PORT"),
	)

	result, err := NewAquilaDb(walletInitStruct).DeleteDocument(&docDelete, url)
	if err != nil {
		t.Error("Something went wrong.", err)
	}
	t.Log("=== Delete Document ==========================")
	t.Logf("%+v", result)
	t.Log("=== Delete Document ==========================")
}
