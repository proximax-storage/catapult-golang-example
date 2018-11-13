# catapult-golang-example

Example codes for Catapult Golang SDK

Details are in https://github.com/proximax-storage/nem2-sdk-go/wiki

# Download

```bash
go get -u https://github.com/proximax-storage/nem2-sdk-go 
go get -u https://github.com/proximax-storage/catapult-golang-example
``` 

# Prerequisite

Config in `$GOPATH/src/github.com/proximax-storage/catapult-golang-example/catapult_config/catapult_config.go`

```go
package catapult_config
const CatapultUrl = "http://privatetest1.proximax.io:3000"
```

# Example Codes
## basic_functions
1. Generate a new Key Pair
2. Create an Address from a given Public Key
3. Create an Account from a given Private Key.

```bash
cd catapult-golang-example/basic_functions
go run main.go
```

## get_accounts_info
1. Simple Account API request

```bash
cd catapult-golang-example/get_accounts_info
go run main.go
```

## get_transactions_info
1. Simple Transaction API request

```bash
cd catapult-golang-example/get_transactions_info
go run main.go
```

## get_mosaic_info
1. Simple Mosaic API request
2. Get information for a set of mosaics.
3. Get readable names for a set of mosaics.
4. Get an array of MosaicInfo from mosaics created under provided namespace.

```bash
cd catapult-golang-example/get_mosaic_info
go run main.go
```

## get_namespace_info
1. Simple Namespace API request
2. Get namespaces an account owns.
3. Get readable names for a set of namespaces.
4. Get an array of NamespaceInfo for a given set of addresses.

```bash
cd catapult-golang-example/get_namespace_info
go run main.go
```

## get_blockchain_info
1. Get BlockInfo for a given block height.
2. Get transactions from a block.
3. Get the current height of the chain.
4. Get the current score of the chain.
5. Get an array of BlockInfo for a given block height and limit.
6. Get the storage information.

```bash
cd catapult-golang-example/get_blockchain_info
go run main.go
```

## get_network_info
1. Simple Network API request

```bash
cd catapult-golang-example/get_network_info
go run main.go
```

# Transaction Examples

## transfer_transaction
1. Transfer 0.000001 XPX

```bash
cd catapult-golang-example/transfer_transaction
go run main.go
```

## register_namespace_transaction
**TODO**
```bash
cd catapult-golang-example/register_namespace_transaction
go run main.go
```

## mosaic_definition_transaction
**TODO**
```bash
cd catapult-golang-example/mosaic_definition_transaction
go run main.go
```

## mosaic_supply_change_transaction
**TODO**
```bash
cd catapult-golang-example/mosaic_supply_change_transaction
go run main.go
```

## modify_multisig_account_transaction
**TODO**
```bash
cd catapult-golang-example/modify_multisig_account_transaction
go run main.go
```

## aggregate_complete_transaction
**TODO**
```bash
cd catapult-golang-example/aggregate_complete_transaction
go run main.go
```

## aggregate_bonded_transactions
**TODO**
```bash
cd catapult-golang-example/aggregate_bonded_transactions
go run main.go
```

## signing_announced_aggregate_bonded_transactions
**TODO**
```bash
cd catapult-golang-example/signing_announced_aggregate_bonded_transactions
go run main.go
```

## websocket
**TODO**
```bash
cd catapult-golang-example/websocket
go run main.go
```
