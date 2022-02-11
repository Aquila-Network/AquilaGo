package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AquilaDbStruct struct {
}

func NewAquilaDb() *AquilaDbStruct {
	return &AquilaDbStruct{}
}

// /db/create
func (a *AquilaDbStruct) CreateDatabase(createDb *CreateDbRequestStruct, url string) (*CreateAquilaResponsStruct, error) {

	var responseAquilaDb *CreateAquilaResponsStruct
	data, err := json.Marshal(createDb)

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
	fmt.Println(string(body)) // write response in the console

	return responseAquilaDb, nil
}

// ???
func (a *AquilaDbStruct) SignDocument() {}

// Send vectors to Aquila DB
// Response will be an array of ids
// /db/doc/insert
func (a *AquilaDbStruct) InsertDocument(docInsert *DocInsertRequestStruct, url string) (*DocInsertResponseStruct, error) {

	var docInsertResponse *DocInsertResponseStruct

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(docInsert)
	if err != nil {
		return docInsertResponse, err
	}

	req, err := json.Marshal(docInsert)
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

	fmt.Println(string(body)) // will write response in the console
	json.Unmarshal(body, &docInsertResponse)

	return docInsertResponse, nil
}

// Delelete Document
// Deleted ids in response
// /db/doc/delete
func (a *AquilaDbStruct) DeleteDocument(docDelete *DocDeleteRequestStruct, url string) (*DocDeleteResponseStruct, error) {

	var docDeleteResponse *DocDeleteResponseStruct

	data, err := json.Marshal(docDelete)

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
func (a *AquilaDbStruct) SearchKDocument(searchBody *SearchAquilaDbRequestStruct, url string) (*DocSearchResponseStruct, error) {
	var docSearchResponse *DocSearchResponseStruct

	// get request
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(searchBody)
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

	/*
		// post request
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(searchBody)
		if err != nil {
			log.Fatal(err)
		}

		req, _ := json.Marshal(searchBody)

		resp, err := http.Post(
			// createURL,
			"https://httpbin.org/post",
			"application/json",
			bytes.NewBuffer(req),
		)
		if err != nil {
			print(err)
		}
		defer resp.Body.Close()
	*/

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return docSearchResponse, err
	}

	// fmt.Println(string(body)) // will write response in the console

	json.Unmarshal(body, &docSearchResponse)

	return docSearchResponse, nil
}
