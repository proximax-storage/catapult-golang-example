package main

import (
	"context"
	"encoding/json"
	"fmt"
	c "github.com/proximax-storage/catapult-golang-example/catapult_config"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"math/big"
)

func main() {

	//--------------------------------------------------
	// 1. Get BlockInfo for a given block height.
	//--------------------------------------------------
	fmt.Println("1. Get BlockInfo for a given block height.")
	conf, err := sdk.NewConfig(c.CatapultUrl, c.NetworkType)
	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	height := big.NewInt(1)

	getBlockByHeight, resp, err := client.Blockchain.GetBlockByHeight(context.Background(), height)
	if err != nil {
		fmt.Printf("Blockchain.GetBlockByHeight returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	getBlockByHeightJson, _ := json.MarshalIndent(getBlockByHeight, "", " ")
	fmt.Printf("%s\n\n", string(getBlockByHeightJson))

	//--------------------------------------------------
	// 2. Get transactions from a block.
	//--------------------------------------------------
	fmt.Println("2. Get transactions from a block.")

	getBlockTransactions, resp, err := client.Blockchain.GetBlockTransactions(context.Background(), height)
	if err != nil {
		fmt.Printf("Blockchain.GetBlockTransactions returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", getBlockTransactions)

	//--------------------------------------------------
	// 3. Get the current height of the chain.
	//--------------------------------------------------
	fmt.Println("3. Get the current height of the chain.")

	getBlockchainHeight, resp, err := client.Blockchain.GetBlockchainHeight(context.Background())
	if err != nil {
		fmt.Printf("Blockchain.GetBlockchainHeight returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", getBlockchainHeight)

	//--------------------------------------------------
	// 4. Get the current score of the chain.
	//--------------------------------------------------
	fmt.Println("4. Get the current score of the chain.")

	getBlockchainScore, resp, err := client.Blockchain.GetBlockchainScore(context.Background())
	if err != nil {
		fmt.Printf("Blockchain.GetBlockchainScore returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", getBlockchainScore)

	//--------------------------------------------------
	// 5. Get an array of BlockInfo for a given block height and limit.
	// TODO: Not Implemented?
	//--------------------------------------------------
	//fmt.Println("5. Get an array of BlockInfo for a given block height and limit.")
	//limit :=  big.NewInt(100)
	//
	//getBlockchainInfo, resp, err := client.Blockchain.GetBlockchainInfo(context.Background(), height, limit)
	//if err != nil {
	//	fmt.Printf("Blockchain.GetBlockchainInfo returned error: %s", err)
	//	return
	//}
	//fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	//fmt.Printf("%s\n\n", getBlockchainInfo)

	//--------------------------------------------------
	// 6. Get the storage information.
	//--------------------------------------------------
	fmt.Println("6. Get the storage information.")

	getBlockchainStorage, resp, err := client.Blockchain.GetBlockchainStorage(context.Background())
	if err != nil {
		fmt.Printf("Blockchain.GetBlockchainStorage returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	getBlockchainStorageJson, _ := json.MarshalIndent(getBlockchainStorage, "", " ")
	fmt.Printf("%s\n\n", getBlockchainStorageJson)
}
