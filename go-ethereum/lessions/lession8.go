package lessions

import (
	"context"
	"fmt"
	token "go-ethereum/erc20"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ReadTokenBalance(client *ethclient.Client) {
	tokenAddress := common.HexToAddress("0x0eb97F0f7CeB62E4286C4871e925EC47FAEd7Ab6")
	instance, _ := token.NewToken(tokenAddress, client)

	account := common.HexToAddress("0xea8093f01def09359c4a2d51699a9085f1159fa9")
	tokenBla, _ := instance.BalanceOf(&bind.CallOpts{}, account)
	fmt.Printf("Token Balance: %v\n", tokenBla)

	totalsupply, _ := instance.TotalSupply(&bind.CallOpts{})
	fmt.Printf("totalsupply: %v\n", totalsupply)

	pri, _ := crypto.HexToECDSA("4f360b3779963dd31921228b734d4153d682b289b192872883bec6551a529a0c")

	chainId, _ := client.NetworkID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(pri, chainId)

	tx, _ := instance.Approve(auth, account, big.NewInt(100000000000000000))
	fmt.Printf("tx: %v\n", tx)

}

// TransferUsingBinding 使用生成的 binding 调用 transfer(to, value)
func TransferUsingBinding(client *ethclient.Client) error {

	// 2) 私钥与 auth
	priv, err := crypto.HexToECDSA("4f360b3779963dd31921228b734d4153d682b289b192872883bec6551a529a0c")
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("network id: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}
	// 可选：覆盖 gas price / gas limit / nonce 等
	// auth.GasLimit = uint64(100000)
	// auth.GasPrice = big.NewInt(20_000_000_000)

	// 3) 绑定合约
	tokenAddr := common.HexToAddress("0x0eb97F0f7CeB62E4286C4871e925EC47FAEd7Ab6")
	tokenInstance, err := token.NewToken(tokenAddr, client)
	if err != nil {
		return fmt.Errorf("new token binding: %w", err)
	}

	// 4) 调用 Transfer
	to := common.HexToAddress("0xea8093f01def09359c4a2d51699a9085f1159fa9")
	value := new(big.Int)
	value.SetString("100000000000000000000", 10)
	tx, err := tokenInstance.Transfer(auth, to, value)
	if err != nil {
		return fmt.Errorf("transfer failed: %w", err)
	}

	fmt.Printf("sent transfer tx: %s\n", tx.Hash().Hex())

	// 5) 可选：等待确认（示例轮询等待）
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}
	if receipt.Status != 1 {
		return fmt.Errorf("tx reverted, status=%v", receipt.Status)
	}
	fmt.Printf("transfer confirmed in block %d\n", receipt.BlockNumber.Uint64())
	return nil
}
