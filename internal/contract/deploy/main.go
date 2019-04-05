package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/islishude/red-package-dapp/internal/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("请输入RPC地址")
	var rpcapi string
	fmt.Scan(&rpcapi)
	if rpcapi == "" {
		rpcapi = "http://127.0.0.1:8545"
	}

	client, err := ethclient.Dial(rpcapi)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fmt.Println("请输入私钥")
	var prikey string
	fmt.Scan(&prikey)

	privateKey, err := crypto.HexToECDSA(prikey)
	if err != nil {
		log.Fatal(err)
	}

	publicKeyECDSA := privateKey.Public().(*ecdsa.PublicKey)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.From = crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("创建者账户地址", auth.From.Hex())

	address, tx, _, err := contract.DeployContract(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sucessful deploy!")
	fmt.Println("Contract address：", address.Hex())
	fmt.Println("Trx Hash：", tx.Hash().Hex())
}
