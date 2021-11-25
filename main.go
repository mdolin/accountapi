package main

import (
	"accountapi/client"
	"accountapi/model"
	"encoding/json"
	"fmt"
	"log"
)

// // http.Client to reuse throughout methods.
// func HttpClient() *http.Client {
// 	client := &http.Client{Timeout: 10 * time.Second}
// 	return client
// }

func fetchData(url string) {
	client := client.CreateClient()
	getData, err := client.Get(url)
	fmt.Println(err)

	// Pretty-printing
	var responseObject model.AccountData
	json.Unmarshal(getData, &responseObject)
	f, _ := json.MarshalIndent(responseObject, "", " ")
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
	toDelete := url + "/" + id + "?version=0"

	client := client.CreateClient()
	d, err := client.Delete(toDelete)

	fmt.Println((err))
	deletedData := string(d)
	log.Printf(deletedData)
}

func main() {
	url := "http://0.0.0.0:8080/v1/organisation/accounts"

	fetchData(url)

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

	// id := "123e4567-e89b-12d3-a456-426614174007"

	// deleteData(url, id)

}