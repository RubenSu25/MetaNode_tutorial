package lessions

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GenerateKey() {
	privateKey, _ := crypto.GenerateKey()
	fmt.Printf("privateKey: %v\n", privateKey)

	// privateKey1,_  := crypto.HexToECDSA("")
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Printf("hexutil.Encode(privateKeyBytes)[2:]: %v\n", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("hexutil.Encode(publicKeyBytes): %v\n", hexutil.Encode(publicKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("address: %v\n", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Printf("hexutil.Encode(hash.Sum(nil)[:]): %v\n", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Printf("hexutil.Encode(hash.Sum(nil)[12:]): %v\n", hexutil.Encode(hash.Sum(nil)[12:]))
}
