## Description
This is Form3 Take Home Exercise done as part of the hiring process written in Go Programming language. The goal was to write a client library in Go to access fake account API which is provided as a Docker container, implement the Create, Fetch, and Delete operations on the accounts resource with all testing.

This was my first project in the Go programming language, so I was using different resources to learn Go and write this client library. Resources are listed below in the [Useful resources](#useful-resources) section.

## Main bits of the project
* The HTTP client
* Create, Fetch, and Delete operations
* Unit tests
* Integration tests

## Structure of the project
```
.
├── account
│   ├── account_create.go
│   ├── account_create_test.go
│   ├── account_delete.go
│   ├── account_delete_test.go
│   ├── account_fetch.go
│   └── account_fetch_test.go
├── account_data.json
├── client
│   ├── client.go
│   └── client_test.go
├── docker-compose.yml
├── go.mod
├── LICENSE
├── main.go
├── model
│   └── models.go
├── README.md
├── scripts
│   └── db
│       └── 10-init.sql
└── tests
    └── integrations
        └── account_test.go

```

## Requirements
To run the project and tests you will need
* [Go Programming language](https://go.dev/doc/install)
* [Docker](https://www.docker.com/get-started)

## Examples
Example for Create data assumes that data for account creating resides in the JSON format in the file located in the same directory as the main function. In this project there is account_data.json

### Example for Fetch data

```Go
package main

import (
	"accountapi/account"
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

func main() {
	namespace := "/v1/organisation/accounts"
	url := "http://0.0.0.0:8080"
	id := "123e4567-e89b-12d3-a456-426614174123"

	FetchData(url+namespace, id)
}
```

### Example for Create data
```Go
package main

import (
	"accountapi/account"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func CreateData(url string, data []byte) {
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

func OpenFile() []byte {
	jsonFile, err := os.Open("account_data.json")

	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func main() {
	namespace := "/v1/organisation/accounts"
	url := "http://0.0.0.0:8080"

	var data = OpenFile()
	CreateData(url+namespace, data)
}
```

### Example for Delete data
```Go
package main

import (
	"accountapi/account"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

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

func main() {
	namespace := "/v1/organisation/accounts"
	url := "http://0.0.0.0:8080"
	id := "123e4567-e89b-12d3-a456-426614174123"

	DeleteData(url+namespace, id)

}
```

## Testing
To run client tests

```
cd account
go test -v
```

To run HTTP client tests

```
cd client
go test -v
```

To run integration tests
```
cd tests/integration
go test -v
```

## Useful resources
* https://go.dev/doc/tutorial/
* https://pkg.go.dev/std
* https://www.practical-go-lessons.com/
* https://blog.logrocket.com/making-http-requests-in-go/
* https://blog.alexellis.io/golang-writing-unit-tests/
* https://mholt.github.io/json-to-go/
