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
