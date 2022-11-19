package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("Hello World!")
	url := "https://cloudflare-eth.com"
	client, err := ethclient.Dial(url)
	check_err(err)
	fmt.Println("[*] Connected to", url)
	var done [10]int
	var found []addresses
	for i := 0; i < 10; i++ {
		go worker(client, i, &done, &found)
	}
	go doneWatcher(&done, &found)
	for {
	}

}

type addresses struct {
	privateKey ecdsa.PrivateKey
	address    string
	balance    big.Int
}

func worker(client *ethclient.Client, id int, done *[10]int, found *[]addresses) {
	counter := 0
	for {
		address, privateKey := generateAddress()
		if balance, _ := client.BalanceAt(context.Background(), common.HexToAddress(address), nil); balance.Cmp(big.NewInt(0)) == 1 {
			*found = append(*found, addresses{*privateKey, address, *balance})
			fmt.Println(privateKey, balance)
		}
		counter++
		if counter%10 == 0 {
			done[id] = counter
		}
	}
}

func doneWatcher(done *[10]int, found *[]addresses) {
	oldDone := *done
	for {
		// fmt.Println(done, oldDone)
		if oldDone != *done {
			fmt.Printf("\x1bc")
			for worker := range *done {
				fmt.Printf("Worker [%d] Tested %d Addresses\n", worker, done[worker])
			}
			fmt.Println("Found", *found)
			// fmt.Println(done)
			oldDone = *done
		}
	}
}

func generateAddress() (address string, privateKey *ecdsa.PrivateKey) {

	privateKey, err := crypto.GenerateKey()
	check_err(err)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return
}

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
