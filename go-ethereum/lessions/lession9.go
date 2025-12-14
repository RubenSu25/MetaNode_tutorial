package lessions

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubscribeNewHead() {
	client, _ := ethclient.Dial("wss://sepolia.infura.io/ws/v3/0d29c1edb43846bc968efc86205dca62")
	defer client.Close()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Failed to subscribe to new heads: %v", err)
	}

	// 接收并处理区块头
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case header := <-headers:
			fmt.Printf("New block header: %v\n", header.TxHash.Hex())
		}
	}

}
