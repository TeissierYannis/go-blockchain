# Build
Build the exe
```
go build
```

# Usage

```
[EXECUTABLE]
    getbalance -address ADDRESS - get the balance for an address
    createblockchain -address ADDRESS creates a blockchain and sends genesis reward to address
    printchain - Prints the blocks in the chain
    send -from FROM -to TO -amount AMOUNT - Send amount of coins
    createwallet -  Creates a new Wallet
    listaddresses - Lists the addresses in our wallet file
    reindexutxo - Rebuilds the UTXO set
```