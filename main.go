package main

import (
	"accountapi/account"
	"accountapi/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func FetchData(url string, id string) {
	var request account.RequestFetch
	request.Host = url
	request.AccountID = id

	response, err := account.AccountFetch(&request)

	if err != nil {
		log.Println("Error: ", err)
	}

	f, _ := json.MarshalIndent(response, "", " ")
	log.Println(string(f))
}

func CreateData(url string, data *model.AccountData) {
	var request account.RequestCreate
	request.Host = url
	request.Data = data

	response, err := account.AccountCreate(&request)

	if err != nil {
		log.Println("Error: ", err)
	}

	f, _ := json.MarshalIndent(response, "", " ")
	log.Println(string(f))
}

func DeleteData(url string, id string) {
	var request account.RequestDelete
	request.Host = url
	request.AccountID = id
	request.Version = "?version=0"

	err := account.AccountDelete(&request)

	if err != nil {
		log.Println("Error: ", err)
	}
}

func OpenFile() *model.AccountData {
	jsonFile, err := os.Open("account_data.json")

	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var accountData model.AccountData
	er := json.Unmarshal([]byte(byteValue), &accountData)
	if er != nil {
		log.Println(er)
	}

	return &accountData
}

func main() {
	namespace := "/v1/organisation/accounts"
	url := "http://0.0.0.0:8080"
	id := "123e4567-e89b-12d3-a456-426614174123"

	FetchData(url+namespace, id)

	// var data = OpenFile()
	// CreateData(url+namespace, data)

	// DeleteData(url+namespace, id)

}
