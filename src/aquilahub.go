package src

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AquilaHubStruct struct {
	Wallet WalletStruct
}

func NewAquilaHub(wallet WalletStruct) *AquilaHubStruct {
	return &AquilaHubStruct{
		Wallet: wallet,
	}
}

// /prepare
func (a *AquilaHubStruct) CreateDatabase(createDb *DataStructCreateDb, url string) (*CreateAquilaHubResponsStruct, error) {

	createDbRequest := &CreateDbRequestStruct{
		Data:      *createDb,
		Signature: a.Wallet.SecretKey,
	}
	// fmt.Println("============================")
	// fmt.Printf("%+v", createDbRequest)

	var responseAquilaDb *CreateAquilaHubResponsStruct
	data, err := json.Marshal(createDbRequest)

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return responseAquilaDb, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseAquilaDb, err
	}

	json.Unmarshal(body, &responseAquilaDb)
	// fmt.Println("============================")
	// fmt.Println(string(body)) // write response in the console

	return responseAquilaDb, nil
}

// Send text array to Aquila hub.
// Response will be an array of vectors.
// /compress
func (d *AquilaHubStruct) CompressDocument(a *AquilaHubRequestStruct, url string) (*AquilaHubResponseStruct, error) {

	var aquilaHubResponse *AquilaHubResponseStruct

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(a)
	if err != nil {
		return aquilaHubResponse, err
	}

	req, _ := json.Marshal(a)

	resp, err := http.Post(
		url,
		// "https://httpbin.org/post", // for debugging
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		return aquilaHubResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return aquilaHubResponse, err
	}

	// fmt.Println(string(body)) // will write response in the console

	json.Unmarshal(body, &aquilaHubResponse)

	return aquilaHubResponse, nil
}
