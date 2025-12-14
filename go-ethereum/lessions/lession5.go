package lessions

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransferETH(client *ethclient.Client) {

	privateKey, _ := crypto.HexToECDSA("4f360b3779963dd31921228b734d4153d682b289b192872883bec6551a529a0c")
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)

	gasPrice, _ := client.SuggestGasPrice(context.Background())
	gasLimit := uint64(21000)

	toAddress := common.HexToAddress("0xea8093f01def09359c4a2d51699a9085f1159fa9")

	amount := big.NewInt(1000000000000000)

	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	chainID, _ := client.NetworkID(context.Background())

	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
}
