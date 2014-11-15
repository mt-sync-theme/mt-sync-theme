package main

import (
	"testing"
)

func TestClientOptions(t *testing.T) {
	opts := clientOptions{
		cmdOptions: &cmdOptions{
			OptEndpoint:   "http://www.example.com/mt/mt-data-api.cgi",
			OptApiVersion: "1",
			OptClientId:   "go-test",
			OptUsername:   "Melody",
		},
		PasswordData: "password",
	}

	if opts.Endpoint() != "http://www.example.com/mt/mt-data-api.cgi" {
		t.Errorf("got %q", opts.Endpoint())
	}

	if opts.ApiVersion() != "1" {
		t.Errorf("got %q", opts.ApiVersion())
	}

	if opts.ClientId() != "go-test" {
		t.Errorf("got %q", opts.ClientId())
	}

	if opts.Username() != "Melody" {
		t.Errorf("got %q", opts.Username())
	}

	if opts.Password() != "password" {
		t.Errorf("got %q", opts.Password())
	}
}
