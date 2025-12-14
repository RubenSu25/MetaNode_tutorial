package lessions

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func TransferToken(client *ethclient.Client) {

	privateKey, _ := crypto.HexToECDSA("4f360b3779963dd31921228b734d4153d682b289b192872883bec6551a529a0c")
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	chainID, _ := client.NetworkID(context.Background())
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)

	fmt.Printf("fromAddress.Hex(): %v\n", fromAddress.Hex())

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("gasPrice : %v", gasPrice)
	}

	datransferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(datransferFnSignature)
	methodId := hash.Sum(nil)[:4]

	toAddress := common.HexToAddress("0xea8093f01def09359c4a2d51699a9085f1159fa9")
	tokenAddress := common.HexToAddress("0x0eb97F0f7CeB62E4286C4871e925EC47FAEd7Ab6")

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal("gasLimit : %v", gasLimit)
	}

	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	sifneder, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	error := client.SendTransaction(context.Background(), sifneder)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("tx.Hash().Hex(): %v\n", tx.Hash().Hex())
}
