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
