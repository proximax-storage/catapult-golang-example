package main

import (
	"context"
	"fmt"
	c "github.com/proximax-storage/catapult-golang-example/catapult_config"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"time"
)

func main() {

	//--------------------------------------------------
	// 1. Transfer 0.000001 XPX
	//--------------------------------------------------
	fmt.Println("1. Transfer 0.000001 XPX")

	// Testnet config default
	conf, err := sdk.NewConfig(c.CatapultUrl, c.NetworkType)
	if err != nil {
		panic(err)
	}

	// Create an account from a private key
	acc, err := sdk.NewAccountFromPrivateKey(c.SenderPrivateKey, c.NetworkType)

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Create a new transfer type transaction
	ttx, err := sdk.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The address of the recipient account.
		sdk.NewAddress(c.RecipientAddress, c.NetworkType),
		// The array of mosaic to be sent - XPX 0.000001
		[]*sdk.Mosaic{sdk.Xem(1)},
		// The transaction message of 1024 bytes.
		sdk.NewPlainMessage(""),
		c.NetworkType,
	)

	// Sign transaction
	stx, err := acc.Sign(ttx)
	if err != nil {
		panic(fmt.Errorf("TransaferTransaction signing returned error: %s", err))
	}

	// Announce transaction
	restTx, err := client.Transaction.Announce(context.Background(), stx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", restTx)
	fmt.Printf("Hash: \t\t%v\n", stx.Hash)
	fmt.Printf("Signer: \t%X\n", acc.KeyPair.PublicKey.Raw)
}
