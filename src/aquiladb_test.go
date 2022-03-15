package src

import (
	"io/ioutil"
	"testing"
)

func TestCreateDatabase(t *testing.T) {
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

	result, err := NewAquilaDb(walletInitStruct).CreateDatabase(&createAquilaDb, "http://localhost:5001/db/create")
	if err != nil {
		t.Error("Something went wrong.", err)
	}
	t.Log("=== Create database ==========================")
	t.Logf("%+v", result)
	t.Log("=== Create database ==========================")
}

func TestInsertDocument(t *testing.T) {

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

	// wallet
	priv, err := ioutil.ReadFile("/home/dev/aquilax/ossl/private_unencrypted.pem")
	if err != nil {
		t.Error(err)
	}
	walletInitStruct := NewWallet(string(priv[:]))
	walletSign, err := walletInitStruct.CreateSignatureWallet(docInsert)
	if err != nil {
		t.Error(err)
	}
	walletInitStruct.SecretKey = walletSign

	result, err := NewAquilaDb(walletInitStruct).InsertDocument(&docInsert, "http://localhost:5001/db/doc/insert")
	if err != nil {
		t.Error("Something went wrong.", err.Error())
	}
	t.Log("=== Doc Insert ==========================")
	t.Logf("%+v", result)
	t.Log("=== Doc Insert ==========================")
}

// func TestSearch(t *testing.T) {
// 	matrix := make([][]float64, 1)
// 	matrix[0] = make([]float64, 1)
// 	matrix[0] = []float64{
// 		-0.01806008443236351, -0.17380790412425995, 0.03992759436368942, 0.43514639139175415,
// 	}
// 	searchBody := DataSearchStruct{
// 		Matrix:       matrix,
// 		K:            10,
// 		R:            0,
// 		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
// 	}

// 	// wallet
// 	priv, err := ioutil.ReadFile("/home/dev/aquilax/ossl/private_unencrypted.pem")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	walletInitStruct := NewWallet(string(priv[:]))
// 	walletSign, err := walletInitStruct.CreateSignatureWallet(searchBody)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	walletInitStruct.SecretKey = walletSign

// 	result, err := NewAquilaDb(walletInitStruct).SearchKDocument(&searchBody, "http://localhost:5001/db/search")
// 	if err != nil {
// 		t.Error("Something went wrong.", err)
// 	}
// 	t.Log("=== Doc Search ==========================")
// 	t.Logf("%+v", result)
// 	t.Log("=== Doc Search ==========================")
// }

func TestDeleteDocument(t *testing.T) {
	var docDelete = DeleteDataStruct{
		Ids: []string{
			"3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
			"BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ",
		},
		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
	}

	// wallet
	priv, err := ioutil.ReadFile("/home/dev/aquilax/ossl/private_unencrypted.pem")
	if err != nil {
		t.Error(err)
	}
	walletInitStruct := NewWallet(string(priv[:]))
	walletSign, err := walletInitStruct.CreateSignatureWallet(docDelete)
	if err != nil {
		t.Error(err)
	}
	walletInitStruct.SecretKey = walletSign

	result, err := NewAquilaDb(walletInitStruct).DeleteDocument(&docDelete, "http://localhost:5001/db/doc/delete")
	if err != nil {
		t.Error("Something went wrong.", err)
	}
	t.Log("=== Delete Document ==========================")
	t.Logf("%+v", result)
	t.Log("=== Delete Document ==========================")
}
