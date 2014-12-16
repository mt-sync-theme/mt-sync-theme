# Movable Type Data API Library For The Go Programming Language

This library of the golang helps you to access to Movable Type's DataAPI.

## Note

* This is a private project.
    * All the responsibility for this project is in Taku AMANO.
* This is not a part of the Movable Type.

## How to use

### Retrieve public entries

```go
package main

import (
	"fmt"
	"github.com/usualoma/mt-data-api-sdk-go"
)

type Entry struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type EntriesResult struct {
	dataapi.Result
	TotalResults int     `json:"totalResults"`
	Items        []Entry `json:"items"`
}

func main() {
	client := dataapi.NewClient(dataapi.ClientOptionsStruct{
		OptEndpoint:   "http://example.com/path/to/mt/mt-data-api.cgi",
		OptApiVersion: "1",
		OptClientId:   "go",
	})

	result := EntriesResult{}
	err := client.SendRequest(
		"GET",
		"/sites/1/entries",
		&dataapi.RequestParameters{
			"searchFields": "keywords",
			"search":       "golang",
		},
		&result)

	if err != nil {
		panic(err)
	}

	if result.Error != nil {
		panic(result.Error.Message)
	}

	fmt.Printf("total: %d\n", result.TotalResults)
	for _, e := range result.Items {
		fmt.Printf("%d: %s\n", e.Id, e.Title)
	}
}
```

### Post a new entry

```go
package main

import (
	"fmt"
	"github.com/usualoma/mt-data-api-sdk-go"
)

type Entry struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type EntryResult struct {
	dataapi.Result
	Entry
}

func main() {
	client := dataapi.NewClient(dataapi.ClientOptionsStruct{
		OptEndpoint:   "http://example.com/path/to/mt/mt-data-api.cgi",
		OptApiVersion: "1",
		OptClientId:   "go",
		OptUsername:   "Melody",
		OptPassword:   "password",
	})

	result := EntryResult{}
	err := client.SendRequest(
		"POST",
		"/sites/1/entries",
		&dataapi.RequestParameters{
			"entry": Entry{
				Title: "Hello golang",
			},
		},
		&result)

	if err != nil {
		panic(err)
	}

	if result.Error != nil {
		panic(result.Error.Message)
	}

	fmt.Printf("%d: %s\n", result.Id, result.Title)
}
```

## LICENSE

Copyright (c) 2014 Taku AMANO

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
