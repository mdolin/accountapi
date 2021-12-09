package main

import (
	"accountapi/account"
	"accountapi/client"
	"encoding/json"
	"fmt"
	"log"
)

func fetchData(url string, id string) {
	var request account.RequestFetch
	request.Host = url
	request.AccountID = id

	response, err := account.AccountFetch(&request)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	f, _ := json.MarshalIndent(response, "", " ")
	log.Println(string(f))
}

func createData(url string, data string) {
	// Convert to byte
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	postBody, err := json.Marshal(result)

	client := client.CreateClient()
	p, err := client.Post(url, postBody)

	fmt.Println(err)
	postData := string(p)
	log.Println(postData)
}

func deleteData(url string, id string) {
	var request account.RequestDelete
	request.Host = url
	request.AccountID = id
	request.Version = "?version=0"

	response, err := account.AccountDelete(&request)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if response == 204 {
		log.Println(response, "Request succeeded")
	}

}

func main() {
	namespace := "/v1/organisation/accounts"
	url := "http://0.0.0.0:8080"
	id := "123e4567-e89b-12d3-a456-426614174111"

	// fetchData(url+namespace, id)

	// data := `{
	// 	"data": {
	// 	  "id": "123e4567-e89b-12d3-a456-426614174007",
	// 	  "organisation_id": "123e4567-e89b-12d3-a456-426614174007",
	// 	  "type": "accounts",
	// 	  "attributes": {
	// 		 "country": "GB",
	// 		  "base_currency": "GBP",
	// 		  "bank_id": "400302",
	// 		  "bank_id_code": "GBDSC",
	// 		  "account_number": "10000004",
	// 		  "customer_id": "234",
	// 		  "iban": "GB28NWBK40030212764204",
	// 		  "bic": "NWBKGB42",
	// 		  "account_classification": "Personal",
	// 		  "name": ["Burek"]
	// 	  }
	// 	}
	//   }`

	// createData(url, data)

	deleteData(url+namespace, id)

}
