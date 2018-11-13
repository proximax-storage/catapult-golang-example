package main

import (
	"encoding/json"
	"fmt"
	c "github.com/proximax-storage/catapult-golang-example/catapult_config"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"golang.org/x/net/context"
)

func main() {
	//--------------------------------------------------
	// 1. Simple Namespace API request
	//--------------------------------------------------
	fmt.Println("1. Simple Namespace API request")
	conf, err := sdk.NewConfig(c.CatapultUrl, c.NetworkType)
	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Generate Id from namespaceName
	namespaceId, _ := sdk.NewNamespaceIdFromName("prx")

	getMosaicsFromNamespace, resp, err := client.Mosaic.GetMosaicsFromNamespace(context.Background(), namespaceId, nil, 0)
	if err != nil {
		fmt.Printf("Namespace.GetNamespace returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)

	getNamespaceJson, _ := json.MarshalIndent(getMosaicsFromNamespace, "", " ")

	fmt.Printf("%s\n\n", string(getNamespaceJson))

	//--------------------------------------------------
	// 2. Get namespaces an account owns.
	//--------------------------------------------------
	fmt.Println("2. Get namespaces an account owns.")
	add := &sdk.Address{Address: c.SenderPublicKey}

	getNamespacesFromAccount, resp, err := client.Namespace.GetNamespacesFromAccount(context.Background(), add, "", 0)
	if err != nil {
		fmt.Printf("Namespace.GetNamespacesFromAccount returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	getNamespacesFromAccountJson, _ := json.MarshalIndent(getNamespacesFromAccount, "", " ")
	fmt.Printf("%s\n\n", string(getNamespacesFromAccountJson))

	//--------------------------------------------------
	// 3. Get readable names for a set of namespaces.
	//--------------------------------------------------
	fmt.Println("3. Get readable names for a set of namespaces.")
	NamespaceIDs := sdk.NamespaceIds{
		List: []*sdk.NamespaceId{
			namespaceId,
		},
	}

	getNamespaceNames, resp, err := client.Namespace.GetNamespaceNames(context.Background(), NamespaceIDs)
	if err != nil {
		fmt.Printf("Namespace.GetNamespaceNames returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	getNamespaceNamesJson, _ := json.MarshalIndent(getNamespaceNames, "", " ")
	fmt.Printf("%s\n\n", string(getNamespaceNamesJson))

	//--------------------------------------------------
	// 4. Get an array of NamespaceInfo for a given set of addresses.
	//--------------------------------------------------
	fmt.Println("4. Get an array of NamespaceInfo for a given set of addresses.")

	var Addresses []*sdk.Address
	aliceAddress := &sdk.Address{Type: c.NetworkType, Address: c.SenderAddress}
	carolAddress := &sdk.Address{Type: c.NetworkType, Address: c.RecipientAddress}

	Addresses = append(Addresses, aliceAddress)
	Addresses = append(Addresses, carolAddress)

	adds := &sdk.Addresses{
		List: Addresses,
	}

	getNamespacesFromAccounts, resp, err := client.Namespace.GetNamespacesFromAccounts(context.Background(), adds, "", 0)
	if err != nil {
		fmt.Printf("Namespace.getNamespacesFromAccounts returned error: %s", err)
		return
	}
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	getNamespacesFromAccountsJson, _ := json.MarshalIndent(getNamespacesFromAccounts, "", " ")
	fmt.Printf("%s\n\n", string(getNamespacesFromAccountsJson))

}
