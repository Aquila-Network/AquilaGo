package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AquilaHubStruct struct {
}

func NewAquilaHub() *AquilaHubStruct {
	return &AquilaHubStruct{}
}

// /prepare
func (a *AquilaHubStruct) CreateHubDatabase(createDb *CreateDbRequestStruct, url string) (*CreateAquilaResponsStruct, error) {

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
