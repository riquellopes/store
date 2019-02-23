package main

import (
	"os"

	"github.com/riquellopes/store/camunda"
)

func main() {
	zee := camunda.Client{
		Address: os.Getenv("CLIENT_ADDRESS"),
		Port:    os.Getenv("CLIENT_PORT"),
	}

	zee.Connect().Deploy()
}
