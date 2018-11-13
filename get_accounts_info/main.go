package main

import (
	"context"
	"fmt"
	c "github.com/proximax-storage/catapult-golang-example/catapult_config"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
)

func main() {
	//--------------------------------------------------
	// 1. Simple Account API request
	//--------------------------------------------------

	addressOne, _ := sdk.NewAddressFromPublicKey(c.SenderPublicKey, c.NetworkType)

	conf, err := sdk.NewConfig(c.CatapultUrl, c.NetworkType)
	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Get AccountInfo for an account.
	// Param address - A Address struct.
	accountInfo, resp, err := client.Account.GetAccountInfo(context.Background(), addressOne)
	if err != nil {
		fmt.Printf("Account.GetAccountInfo returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf(accountInfo.String())
}
