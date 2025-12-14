package lessions

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ReadTxList(client *ethclient.Client) {
	fmt.Printf("\"=\": %v\n", "=============== 2 ===============")
	block, err := client.BlockByNumber(context.Background(), big.NewInt(9823465))
	if err != nil {
		log.Fatal(block)
	}

	txList := block.Transactions()

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range txList {
		fmt.Printf("tx: %v\n", tx.Hash())

		sender, err := types.Sender(types.NewEIP155Signer(chainID), tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sender.Hex(): %v\n", sender.Hex())

		// if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
		// 	fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		// }

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("receipt.Status: %v\n", receipt.Status)
	}

	fmt.Printf("\"=========================\": %v\n", "=========================")

	hash := common.HexToHash("0xdbc0063ffe1ef061910698e268809c254d7f17265048f8e96631f66701474581")
	if count, err := client.TransactionCount(context.Background(), hash); err == nil {
		for i := uint(0); i < count; i++ {
			tx, err := client.TransactionInBlock(context.Background(), hash, i)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("tx.Hash(): %v\n", tx.Hash().Hex())
		}
	}

	txHash := common.HexToHash("0xc3ade937fc956bd2f1bfa010b1f3c374e7442edf5b8fef50aaf28428e5a3a2e8")
	tx, ispending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx.Hash(): %v\n", tx.Hash())
	fmt.Printf("tx.Hash(): %v\n", tx.Hash().Hex())
	fmt.Printf("ispending: %v\n", ispending)

}
