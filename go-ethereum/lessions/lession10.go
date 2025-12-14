package lessions

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	store "go-ethereum/store"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployStoreWithSolc(client *ethclient.Client) {
	pri, _ := crypto.HexToECDSA("4f360b3779963dd31921228b734d4153d682b289b192872883bec6551a529a0c")
	publicKey := pri.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)
	chainId, _ := client.NetworkID(context.Background())
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)

	auth, _ := bind.NewKeyedTransactorWithChainID(pri, chainId)
	auth.From = fromAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = nil                // in wei
	auth.GasLimit = uint64(3000000) // in units
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice

	tokenAddress, tx, instance, err := store.DeployStore(auth, client, "1.0")

	if err != nil {
		fmt.Printf("DeployStore err: %v\n", err)
	}
	fmt.Printf("tx.Hash().Hex(): %v\n", tx.Hash().Hex())
	fmt.Printf("tokenAddress.Hex(): %v\n", tokenAddress.Hex())

	receipt, _ := bind.WaitMined(auth.Context, client, tx)
	if receipt.Status != 1 {
		log.Fatal("deploy failed")
	}
	version, _ := instance.Version(&bind.CallOpts{})
	fmt.Printf("version: %v\n", version)
}

func DeployStore(client *ethclient.Client) {
	pri, _ := crypto.HexToECDSA("4f360b3779963dd31921228b734d4153d682b289b192872883bec6551a529a0c")
	publicKey := pri.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	gasLimit := uint64(53450)

	parsed, _ := abi.JSON(strings.NewReader(store.StoreABI))
	ctorArgs := parsed.Constructor.Inputs
	packed, _ := ctorArgs.Pack("1.0")

	bin := common.Hex2Bytes(store.StoreBin)

	data := append(bin, packed...)

	tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, data)
	chainId, _ := client.NetworkID(context.Background())
	signer := types.NewEIP155Signer(chainId)
	signedTx, _ := types.SignTx(tx, signer, pri)
	error := client.SendTransaction(context.Background(), signedTx)
	if error != nil {
		log.Fatal("SendTransaction err: %v", error)
	}
	fmt.Printf("tx.Hash().Hex(): %v\n", tx.Hash().Hex())

	// 等待一段时间后再次查询
	time.Sleep(12 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal("TransactionReceipt err: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatal("deploy failed")
	}
	tokenAddress := receipt.ContractAddress
	fmt.Printf("tokenAddress.Hex(): %v\n", tokenAddress.Hex())

	instance, err := store.NewStore(tokenAddress, client)
	if err != nil {
		log.Fatal("NewStore err: %v", err)
	}

	time.Sleep(10 * time.Second)

	version, _ := instance.Version(&bind.CallOpts{})
	fmt.Printf("version: %v\n", version)

}
