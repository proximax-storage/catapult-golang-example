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

	getMosaicsFromNamespace, err := client.Mosaic.GetMosaicsFromNamespace(context.Background(), namespaceId, 0)
	if err != nil {
		fmt.Printf("Namespace.GetNamespace returned error: %s", err)
		return
	}

	getNamespaceJson, _ := json.MarshalIndent(getMosaicsFromNamespace, "", " ")

	fmt.Printf("%s\n\n", string(getNamespaceJson))

	//--------------------------------------------------
	// 2. Get namespaces an account owns.
	//--------------------------------------------------
	fmt.Println("2. Get namespaces an account owns.")
	add := &sdk.Address{Address: c.SenderPublicKey}

	getNamespacesFromAccount, err := client.Namespace.GetNamespacesFromAccount(context.Background(), add, nil, 0)
	if err != nil {
		fmt.Printf("Namespace.GetNamespacesFromAccount returned error: %s", err)
		return
	}
	getNamespacesFromAccountJson, _ := json.MarshalIndent(getNamespacesFromAccount, "", " ")
	fmt.Printf("%s\n\n", string(getNamespacesFromAccountJson))

	//--------------------------------------------------
	// 3. Get readable names for a set of namespaces.
	//--------------------------------------------------
	fmt.Println("3. Get readable names for a set of namespaces.")
	var NamespaceIDs []*sdk.NamespaceId
	namespaceId2, _ := sdk.NewNamespaceIdFromName("nem")
	NamespaceIDs = append(NamespaceIDs, namespaceId2)

	getNamespaceNames, err := client.Namespace.GetNamespaceNames(context.Background(), NamespaceIDs)
	if err != nil {
		fmt.Printf("Namespace.GetNamespaceNames returned error: %s", err)
		return
	}
	getNamespaceNamesJson, _ := json.MarshalIndent(getNamespaceNames, "", " ")
	fmt.Printf("%s\n\n", string(getNamespaceNamesJson))

	//--------------------------------------------------
	// 4. Get an array of NamespaceInfo for a given set of addresses.
	//--------------------------------------------------
	fmt.Println("4. Get an array of NamespaceInfo for a given set of addresses.")

	var Addresses []*sdk.Address
	aliceAddress, _ := sdk.NewAddressFromRaw(c.SenderAddress)
	carolAddress, _ := sdk.NewAddressFromRaw(c.RecipientAddress)

	Addresses = append(Addresses, aliceAddress)
	Addresses = append(Addresses, carolAddress)

	getNamespacesFromAccounts, err := client.Namespace.GetNamespacesFromAccounts(context.Background(), Addresses, nil, 0)
	if err != nil {
		fmt.Printf("Namespace.getNamespacesFromAccounts returned error: %s", err)
		return
	}
	getNamespacesFromAccountsJson, _ := json.MarshalIndent(getNamespacesFromAccounts, "", " ")
	fmt.Printf("%s\n\n", string(getNamespacesFromAccountsJson))
}
