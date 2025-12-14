package lessions

import (
	"context"
	"fmt"
	"go-ethereum/store"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckEvent() {

	client, _ := ethclient.Dial("wss://sepolia.infura.io/ws/v3/0d29c1edb43846bc968efc86205dca62")
	defer client.Close()

	tokenAddress := common.HexToAddress("0x452F69AEC51f89b08B65c225428Fc66cb4F64BE3")

	// block, _ := client.BlockByNumber(context.Background(), nil)
	query := ethereum.FilterQuery{
		// ToBlock:   big.NewInt(),
		Addresses: []common.Address{tokenAddress},
	}

	logs, _ := client.FilterLogs(context.Background(), query)

	contractAbi, err := abi.JSON(strings.NewReader(store.StoreABI))
	if err != nil {
		log.Fatal(err)
	}
	for _, vlog := range logs {
		fmt.Printf("log.Address.Hex(): %v\n", vlog.Address.Hex())
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vlog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("common.BytesToHash(event.key[:]): %v\n", common.BytesToHash(event.Key[:]))
		fmt.Printf("common.BytesToHash(event.value[:]): %v\n", common.BytesToHash(event.Value[:]))

		for _, topic := range vlog.Topics {
			fmt.Printf("topic.Hex(): %v\n", topic.Hex())
		}
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Printf("hash: %v\n", hash.Hex())

	fmt.Printf("\"=========================\": %v\n", "=========================")

	var subLogs = make(chan types.Log)
	sub, _ := client.SubscribeFilterLogs(context.Background(), query, subLogs)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vlog := <-subLogs:
			fmt.Printf("vlog.Address.Hex(): %v\n", vlog.Address.Hex())
			for _, topic := range vlog.Topics {
				fmt.Printf("topic.Hex(): %v\n", topic.Hex())
			}

			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}

			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vlog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("common.BytesToHash(event.key[:]): %v\n", common.BytesToHash(event.Key[:]))
			fmt.Printf("common.BytesToHash(event.value[:]): %v\n", common.BytesToHash(event.Value[:]))
		}
	}

}
