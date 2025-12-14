package lessions

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func TransactionReceipts(client *ethclient.Client) {

	block, _ := client.BlockByNumber(context.Background(), nil)

	blockHash := block.Hash()

	blockNumber := block.Number()
	receipts, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}

	receiptsNumbers, err := client.BlockReceipts(context.Background(),
		rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))

	if err != nil {
		log.Fatal("get receipts by number failed: %s", err)
	}

	for _, receipt := range receipts {
		fmt.Printf("receipt.Status: %v\n", receipt.Status)
		fmt.Printf("receipt.Logs: %v\n", receipt.Logs)
	}

	for _, number := range receiptsNumbers {
		fmt.Printf("number.TxHash.Hex(): %v\n", number.TxHash.Hex())
	}

}
