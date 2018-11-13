package main

import (
	"context"
	"encoding/json"
	"fmt"
	c "github.com/proximax-storage/catapult-golang-example/catapult_config"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
)

func main() {
	//--------------------------------------------------
	// 1. Simple Mosaic API request
	//--------------------------------------------------
	fmt.Println("1. Simple Mosaic API request")
	conf, err := sdk.NewConfig(c.CatapultUrl, c.NetworkType)
	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Create a Mosaic identifier.
	// Example Mosaic identifier: "nem:xem"
	mosaicId, _ := sdk.NewMosaicIdFromName("prx:xpx")

	// Get mosaic information.
	getMosaic, resp, err := client.Mosaic.GetMosaic(context.Background(), mosaicId)
	if err != nil {
		fmt.Printf("Mosaic.GetMosaic returned error: %s", err)
		return
	}
	getMosaicJson, _ := json.MarshalIndent(getMosaic, "", " ")
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(getMosaicJson))

	//--------------------------------------------------
	// 2. Get information for a set of mosaics.
	//--------------------------------------------------
	fmt.Println("2. Get information for a set of mosaics.")
	// Create the Mosaics identifier.
	xpxMosaicId, _ := sdk.NewMosaicIdFromName("prx:xpx")
	xemMosaicId, _ := sdk.NewMosaicIdFromName("nem:xem")

	var mosaics sdk.MosaicIds
	// Append attachment into Mosaics
	mosaics.MosaicIds = append(mosaics.MosaicIds, xpxMosaicId)
	// Append attachment into Mosaics
	mosaics.MosaicIds = append(mosaics.MosaicIds, xemMosaicId)

	getMosaics, resp, err := client.Mosaic.GetMosaics(context.Background(), mosaics)
	if err != nil {
		fmt.Printf("Mosaic.GetMosaics returned error: %s", err)
		return
	}
	getMosaicsJson, _ := json.MarshalIndent(getMosaics, "", " ")
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(getMosaicsJson))

	//--------------------------------------------------
	// 3. Get readable names for a set of mosaics.
	//--------------------------------------------------
	fmt.Println("3. Get readable names for a set of mosaics.")
	var mosaics2 sdk.MosaicIds

	// Append attachment into Mosaics
	mosaics2.MosaicIds = append(mosaics2.MosaicIds, xpxMosaicId)
	// Append attachment into Mosaics
	mosaics2.MosaicIds = append(mosaics2.MosaicIds, xemMosaicId)

	getMosaicNames, resp, err := client.Mosaic.GetMosaicNames(context.Background(), mosaics2)
	if err != nil {
		fmt.Printf("Mosaic.GetMosaicNames returned error: %s", err)
		return
	}
	getMosaicNamesJson, _ := json.MarshalIndent(getMosaicNames, "", " ")
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(getMosaicNamesJson))

	//--------------------------------------------------
	// 4. Get an array of MosaicInfo from mosaics created under provided namespace.
	//--------------------------------------------------
	fmt.Println("4. Get an array of MosaicInfo from mosaics created under provided namespace.")
	namespaceId, _ := sdk.NewNamespaceIdFromName("prx")

	getMosaicsFromNamespace, resp, err := client.Mosaic.GetMosaicsFromNamespace(context.Background(), namespaceId, nil, 0)
	if err != nil {
		fmt.Printf("Mosaic.GetMosaicsFromNamespace returned error: %s", err)
		return
	}

	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	for _, m := range getMosaicsFromNamespace {
		getMosaicsFromNamespaceJson, _ := json.MarshalIndent(m, "", " ")
		fmt.Printf("%s\n\n", string(getMosaicsFromNamespaceJson))
	}
}
