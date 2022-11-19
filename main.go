package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("Hello World!")
	url := "https://cloudflare-eth.com"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[*] Connected to", url)
	_ = client
}
