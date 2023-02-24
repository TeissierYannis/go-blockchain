[![GoDoc](https://godoc.org/github.com/abhishekkr/goshield?status.svg)](https://go.dev/doc/)
[![Go version](https://img.shields.io/badge/Go_version-1.18-blue.svg)]()
[![Type](https://img.shields.io/badge/Topic-Blockchain-blue.svg)]()

# Table of Contents

1.  [Project explanation](#project-explanation)
2. [Tree structure](#tree-structure)
3. [Using docker to mock a network of nodes](#using-docker-to-mock-a-network-of-nodes)
   1. [Step 1 - Build the executable](#using-docker-to-mock-a-network-of-nodes-step1)
   2. [Step 2 - Build docker image](#using-docker-to-mock-a-network-of-nodes-step2)
   3. [Step 3 - Run the docker image](#using-docker-to-mock-a-network-of-nodes-step3)
5. [Build](#build)
5. [Usage](#usage)


# Project explanation <a id="project-explanation"></a>

This is a simple blockchain implementation in Go. It is a learning project and is not intended for production use.

This blockchain includes the following features:
- [X] Proof of work
- [X] Transactions
- [X] Wallets
- [X] UTXO
- [X] Merkle trees
- [X] Network
- [X] Persistence
- [X] Docker
- [ ] Unit tests
- [X] CLI
- [ ] RPC API
- [ ] Web interface
- [ ] Advanced P2P network
- [ ] Advanced consensus algorithm
- [ ] Smart contracts
- [ ] Decentralized exchange
- [ ] Decentralized storage

# Tree structure <a id="tree-structure"></a>

```sh
.
├── README.md
├── bin
│   ├── addresses
│   │   ├── node_address_XX
│   │   ├── node_address_XY
│   ├── tmp
│   │   ├── blocks_XXX
│   │   │   ├── DB_FILES
│   │   ├── wallets_XXX.data
│   ├── blockchain_exe_file
│   ├── docker-compose.yml
│   ├── Dockerfile.main
│   ├── Dockerfile.node
│   ├── Dockerfile.wallet
│   ├── main.sh
│   ├── node.sh
│   ├── wallet.sh
├── blockchain
│   ├── block.go
│   ├── blockchain.go
│   ├── chain_iter.go
│   ├── merkle.go
│   ├── proof.go
│   ├── transaction.go
│   ├── tx.go
│   ├── utxo.go
├── cli
│   ├── cli.go
├── network
│   ├── network.go
│   └── utils.go
├── tmp
│   ├── wallets_XXX.data
│   ├── blocks_XXX
│   │   ├─ DB_FILES
├── wallet
│   ├── utils.go
│   ├── wallet.go
│   ├── wallets.go
```


# Build
To build the executable file run the following command:
```sh
go build
```

# Using docker to mock a network of nodes <a id="using-docker-to-mock-a-network-of-nodes"></a>
## Step 1 - Build the executable file<a id="using-docker-to-mock-a-network-of-nodes-step1"></a>
First, build the go executable.
```sh
go build -o blockchain-linux-arm64  # Or another name but make sure to change the name in the docker-compose.yml file
```
## Step 2 - Build docker image <a id="using-docker-to-mock-a-network-of-nodes-step2"></a>
Build the docker images.
```sh
docker-compose build
```

## Step 3 - Run docker containers <a id="using-docker-to-mock-a-network-of-nodes-step3"></a>
Run the docker containers.
```sh
docker-compose up
```

# Build <a id="build"></a>

To build the executable file run the following command:
```sh
go build -o name_of_executable
```

# Usage <a id="usage"></a>
```sh
[EXECUTABLE]
    getbalance -address ADDRESS - get the balance for an address
    createblockchain -address ADDRESS creates a blockchain and sends genesis reward to address
    printchain - Prints the blocks in the chain
    send -from FROM -to TO -amount AMOUNT - Send amount of coins
    createwallet -  Creates a new Wallet
    listaddresses - Lists the addresses in our wallet file
    reindexutxo - Rebuilds the UTXO set
```




