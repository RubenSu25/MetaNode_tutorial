package lessions

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ReadHeaderAndBlock(client *ethclient.Client) {

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("header.TxHash.String(): %v\n", header.TxHash.String())

	block, err := client.BlockByNumber(context.Background(), big.NewInt(9823465))
	if err != nil {
		log.Fatal(block)
	}
	fmt.Printf("block.Hash(): %v\n", block.Hash())
	fmt.Printf("block.NumberU64(): %v\n", block.NumberU64())

	block1, err := client.BlockByHash(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("block1.Coinbase(): %v\n", block1.Coinbase())
	fmt.Printf("block1.Hash(): %v\n", block1.Hash())

}
