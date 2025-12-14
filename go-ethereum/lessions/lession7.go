package lessions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ReadBalance(client *ethclient.Client) {
	account := common.HexToAddress("0xea8093f01def09359c4a2d51699a9085f1159fa9")
	balance, _ := client.BalanceAt(context.Background(), account, nil)
	fmt.Printf("balance: %v\n", balance)

	balance1, _ := client.BalanceAt(context.Background(), account, big.NewInt(56756787))
	fmt.Printf("balance1: %v\n", balance1)

	balForETF := big.NewFloat(0)
	balForETF.SetString(balance.String())
	ethValue := new(big.Float).Quo(balForETF, big.NewFloat(1e18))
	fmt.Printf("ethValue: %f\n", ethValue)

	pending, _ := client.PendingBalanceAt(context.Background(), account)
	fmt.Printf("pending: %v\n", pending)

}
