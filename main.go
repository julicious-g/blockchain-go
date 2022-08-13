package main

import "blockchain-go/blockchain"

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Gema")
	bc.AddBlock("Send 2 BTC to Gema")

	bc.Print()
}
