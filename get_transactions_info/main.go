package main

import (
	"context"
	"fmt"
	c "github.com/proximax-storage/catapult-golang-example/catapult_config"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
)

func main() {

	//--------------------------------------------------
	// 1. Simple Transaction API request
	//--------------------------------------------------

	conf, err := sdk.NewConfig(c.CatapultUrl, c.NetworkType)
	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Get transaction information of transactionId or transactionHash.
	// Param Hash - transactionId or transactionHash.
	// Example transactionId: "5BE6FB4CEBBDEC00012EF3EA"
	// Example transactionHash: "340272E20D100E24732AB69F6F88A5989AD130BA15B37126ADF2A31F85C7F3F6"
	getTransaction, resp, err := client.Transaction.GetTransaction(context.Background(), "5BE6FB4CEBBDEC00012EF3EA")
	if err != nil {
		fmt.Printf("Transaction.GetTransaction returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", getTransaction.String())
}
