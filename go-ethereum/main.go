package main

import (
	"go-ethereum/lessions"
)

func main() {
	// client, err := ethclient.Dial("https://sepolia.infura.io/v3/0d29c1edb43846bc968efc86205dca62")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("client: %v\n", client)
	// defer client.Close()

	// lessions.ReadHeaderAndBlock(client)
	// lessions.ReadTxList(client)
	// lessions.TransactionReceipts(client)
	// lessions.GenerateKey()
	// lessions.TransferETH(client)
	// lessions.TransferToken(client)
	// lessions.ReadBalance(client)

	// lessions.ReadTokenBalance(client)
	// lessions.TransferUsingBinding(client)
	// lessions.SubscribeNewHead()

	// lessions.DeployStoreWithSolc(client)
	// lessions.DeployStore(client)
	// lessions.UseSolidity(client)

	// lessions.UseAbi(client)
	lessions.CheckEvent()

}
