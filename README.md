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
```bash
cd catapult-golang-example/get_accounts_info
go run main.go
```

## get_accounts_info
```bash
cd catapult-golang-example/get_accounts_info
go run main.go
```

## get_transactions_info
```bash
cd catapult-golang-example/get_transactions_info
go run main.go
```

## get_mosaic_info
```bash
cd catapult-golang-example/get_mosaic_info
go run main.go
```

## get_namespace_info
```bash
cd catapult-golang-example/get_namespace_info
go run main.go
```

## get_blockchain_info
```bash
cd catapult-golang-example/get_blockchain_info
go run main.go
```

## get_network_info
```bash
cd catapult-golang-example/get_network_info
go run main.go
```

## announces_a_transaction
```bash
cd catapult-golang-example/announces_a_transaction
go run main.go
```

## transfer_transaction
```bash
cd catapult-golang-example/transfer_transaction
go run main.go
```

## register_namespace_transaction
```bash
cd catapult-golang-example/register_namespace_transaction
go run main.go
```

## mosaic_definition_transaction
```bash
cd catapult-golang-example/mosaic_definition_transaction
go run main.go
```

## mosaic_supply_change_transaction
```bash
cd catapult-golang-example/mosaic_supply_change_transaction
go run main.go
```

## modify_multisig_account_transaction
```bash
cd catapult-golang-example/modify_multisig_account_transaction
go run main.go
```

## aggregate_complete_transaction
```bash
cd catapult-golang-example/aggregate_complete_transaction
go run main.go
```

## aggregate_bonded_transactions
```bash
cd catapult-golang-example/aggregate_bonded_transactions
go run main.go
```

## signing_announced_aggregate_bonded_transactions
```bash
cd catapult-golang-example/signing_announced_aggregate_bonded_transactions
go run main.go
```

## websocket
```bash
cd catapult-golang-example/websocket
go run main.go
```
