package main

import (
	"context"
	"fmt"
	c "github.com/proximax-storage/catapult-golang-example/catapult_config"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
)

func main() {

	//--------------------------------------------------
	// 1. Simple Network API request
	//--------------------------------------------------
	fmt.Println("1. Simple Network API request")
	conf, err := sdk.NewConfig(c.CatapultUrl, c.NetworkType)
	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	getNetworkType, err := client.Network.GetNetworkType(context.Background())
	if err != nil {
		fmt.Printf("Network.GetNetworkType returned error: %s", err)
		return
	}
	fmt.Printf("%s\n\n", getNetworkType)
}
