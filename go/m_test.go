package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
	"testing"
)

var (
	privateKeyString = os.Getenv("BEVM_DEV_PRIVATE_KEY")
	contractString   = "0x30A0e025BE2bbC80948f60647c48756815b78227"
	userAddress      = "0x8a43286e4B655D1f2Aae80F3F6108e455CC0f01d"
)

func TestMint(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
		return
	}
	privateKeyString = os.Getenv("BEVM_DEV_PRIVATE_KEY")

	client, err := ethclient.Dial("wss://testnet.bevm.io/ws")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(contractString)
	//contractABI, err := abi.JSON(strings.NewReader(abiString))
	privateKey, _ := crypto.HexToECDSA(privateKeyString)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	amount := decimal.NewFromBigInt(big.NewInt(10), 18).BigInt()
	to := common.HexToAddress(userAddress)

	contractInstance, err := NewXbtc(contractAddress, client)
	tx, err := contractInstance.Mint(auth, to, amount)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("tx sent: ", tx.Hash().Hex())
	}
}

func TestBurn(t *testing.T) {
	client, err := ethclient.Dial("wss://testnet.bevm.io/ws")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(contractString)
	//contractABI, err := abi.JSON(strings.NewReader(abiString))
	privateKey, _ := crypto.HexToECDSA(privateKeyString)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	amount := decimal.NewFromBigInt(big.NewInt(10), 18).BigInt()
	to := common.HexToAddress(userAddress)

	contractInstance, err := NewXbtc(contractAddress, client)
	tx, err := contractInstance.Burn(auth, to, amount)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("tx sent: ", tx.Hash().Hex())
	}
}
