package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
	"strings"
)

var (
	client *ethclient.Client
)

func main() {
	client, err := ethclient.Dial("wss://testnet.bevm.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x30A0e025BE2bbC80948f60647c48756815b78227")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	dir, err := os.Getwd()
	abiString, err := os.ReadFile(dir + "/go/abi.json")
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	contractAbi, err := abi.JSON(strings.NewReader(string(abiString)))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			fmt.Println("Log Block Number:", vLog.BlockNumber)
			fmt.Println("Log Index:", vLog.Index)
			fmt.Println("Log Data:", vLog.Data)
			fmt.Println("Log Topic:", vLog.Topics)
			fmt.Println("Log TxIndex:", vLog.TxIndex)
			fmt.Println("Log :", vLog)
			receipt, err := client.TransactionReceipt(context.Background(), vLog.TxHash)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("receipt :", receipt)
			fmt.Println("--------------------------------------------------------------------------------", receipt)
			var transferEvent LogTransfer
			err = contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Println(
				"event Transfer", transferEvent.From.Hex(), transferEvent.To.Hex(), transferEvent.Amount.String(),
			)
			fmt.Println(
				"Transaction Value:", vLog.Topics[0].Hex(), vLog.Topics[1].Hex(), decimal.NewFromBigInt(transferEvent.Amount, -18).String(),
			)

			if vLog.Topics[0].Hex() == "Transfer" {
				transaction, _, err := client.TransactionByHash(context.Background(), vLog.TxHash)
				if err != nil {
					log.Fatal(err)
				}
				amount := transaction.Value()
				log.Println("Transaction Amount: ", amount)
			}
		}
	}
}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
}

type mintArgs struct {
	_To     common.Address
	_Amount *big.Int
}
