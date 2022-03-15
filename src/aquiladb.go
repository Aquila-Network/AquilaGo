package src

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AquilaDbStruct struct {
	Wallet WalletStruct
}

func NewAquilaDb(wallet WalletStruct) *AquilaDbStruct {
	return &AquilaDbStruct{
		Wallet: wallet,
	}
}

// /db/create
func (a *AquilaDbStruct) CreateDatabase(createDb *DataStructCreateDb, url string) (*CreateAquilaResponsStruct, error) {

	var responseAquilaDb *CreateAquilaResponsStruct

	createDbRequest := &CreateDbRequestStruct{
		Data:      *createDb,
		Signature: a.Wallet.SecretKey,
	}
	// fmt.Println("==================================")
	// fmt.Printf("%+v", createDbRequest)

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
	// fmt.Println(string(body)) // write response in the console

	return responseAquilaDb, nil
}

// ???
func (a *AquilaDbStruct) SignDocument() {}

// Send vectors to Aquila DB
// Response will be an array of ids
// /db/doc/insert
func (a *AquilaDbStruct) InsertDocument(docInsert *DatatDocInsertStruct, url string) (*DocInsertResponseStruct, error) {

	var docInsertResponse *DocInsertResponseStruct
	insertDbRequest := &DocInsertRequestStruct{
		Data:      *docInsert,
		Signature: a.Wallet.SecretKey,
	}
	// fmt.Println("=========================")
	// fmt.Print(a.Wallet.SecretKey)

	req, err := json.Marshal(insertDbRequest)
	// fmt.Println("=========================")
	// fmt.Print(string(req[:]))  // write json to console
	if err != nil {
		return docInsertResponse, err
	}

	resp, err := http.Post(
		url,
		// "https://httpbin.org/post", // for debugging
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		return docInsertResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return docInsertResponse, err
	}

	// fmt.Println("================ill write re=========")
	// fmt.Println(string(body)) // wsponse in the console
	json.Unmarshal(body, &docInsertResponse)

	return docInsertResponse, nil
}

// Delelete Document
// Deleted ids in response
// /db/doc/delete
func (a *AquilaDbStruct) DeleteDocument(docDelete *DeleteDataStruct, url string) (*DocDeleteResponseStruct, error) {

	var docDeleteResponse *DocDeleteResponseStruct

	deleteDbRequest := &DocDeleteRequestStruct{
		Data:      *docDelete,
		Signature: a.Wallet.SecretKey,
	}
	// fmt.Printf("%+v", deleteDbRequest)

	data, err := json.Marshal(deleteDbRequest)
	// fmt.Println("=========================")
	// fmt.Print(string(data[:])) // write json to console

	resp, err := http.Post(
		url,
		// "https://httpbin.org/post", // route for debugging
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return docDeleteResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return docDeleteResponse, err
	}

	// fmt.Println(string(body)) // write response in the console
	json.Unmarshal(body, &docDeleteResponse)

	return docDeleteResponse, nil
}

// /db/search
func (a *AquilaDbStruct) SearchKDocument(searchBody *DataSearchStruct, url string) (*DocSearchResponseStruct, error) {
	var docSearchResponse *DocSearchResponseStruct

	searchDbStruct := &SearchAquilaDbRequestStruct{
		Data:      *searchBody,
		Signature: a.Wallet.SecretKey,
	}

	// get request
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(searchDbStruct)
	if err != nil {
		return docSearchResponse, err
	}

	req, err := http.NewRequest(http.MethodGet, url, &buf)
	if err != nil {
		return docSearchResponse, err
	}

	// add header to GET request
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	resp, err := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return docSearchResponse, err
	}

	// fmt.Println(string(body)) // will write response in the console
	json.Unmarshal(body, &docSearchResponse)

	return docSearchResponse, nil
}
