package main

import (
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/crypto"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
)

func main() {

	//--------------------------------------------------
	// 1. Generate a new Key Pair
	//--------------------------------------------------
	fmt.Println("1. Generate a new Key Pair")
	KeyPair, _ := crypto.NewRandomKeyPair()

	fmt.Printf("\tGenerated PrivateKey:\t%x\n", KeyPair.PrivateKey.Raw)
	fmt.Printf("\tGenerated PublicKey:\t%x\n\n", KeyPair.PublicKey.Raw)

	//--------------------------------------------------
	// 2. Create an Address from a given public key.
	//--------------------------------------------------
	fmt.Println("2. Create an Address from a given public key.")
	publicKey := string(KeyPair.PublicKey.String())

	// Generate a NEM address with network set to MijinTest (144)
	Address, _ := sdk.NewAddressFromPublicKey(publicKey, sdk.MijinTest)

	fmt.Printf("\tGenerated Address:\t%v\n", Address.Address)
	fmt.Printf("\tUsed NetworkType:\t%s\n\n", Address.Type)

	//--------------------------------------------------
	// 3. Create an Account from a given private key.
	//--------------------------------------------------
	fmt.Println("3. Create an Account from a given private key.")
	privateKey := string(KeyPair.PrivateKey.String())

	newAccount, _ := sdk.NewAccountFromPrivateKey(privateKey, sdk.MijinTest)

	fmt.Printf("\tPrivateKey:\t%x\n", newAccount.KeyPair.PrivateKey.Raw)
	fmt.Printf("\tPublicKey:\t%x\n", newAccount.KeyPair.PublicKey.Raw)
	fmt.Printf("\tAddress:\t%v\n\n", newAccount.Address.Pretty())
}
