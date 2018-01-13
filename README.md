# Digital - A basic blockchain

**caution**: *Digital* is in an early stage and it is still under development

## Introduction
**Blockchain** is one of the hot topics of the moment. It seems that it has the potential to change the way that the world approaches distributed ledgers but, what is a blockchain?

A Blockchain is a distributed database system that acts as an “open ledger” to store and manage transactions. The key idea behind is that it is formed by blocks. These blocks contain details about the trasnactions processed at some point and a link to the block which was processed just before. The way of linking one block with the previous one (creating a "block"-"chain") is adding the *hash* of the previous block to the new block and its own *hash*. This makes it impossible for anyone to alter information about the records retrospectively which makes blockchain **immutable** — information remains in the same state for as long as the network exists.

This technology was first introduced in the [Bitcoin whitepaper](https://bitcoin.org/bitcoin.pdf), in 2009 by Satoshi Nakamoto.

After some hype on the cryptocurrencies some people saw the value on the technology that enables Bitcoin,  *Blockchain* a public ledger shared by all the entities in wich one can trust.

**Digital**

## Objectives
Digital is an implementation of a simple blockchain in Golang based on bitcoin. This repository contains the "core" library of the project. That means that you will not find here the necessary code for setting up a node in a network. This constrain has been intentionally because the idea is to help others to learn the key concepts behind blockchain and not get a production ready implementation.

The main objectives of *Digital* are:
- Become the entry point for **understanding the key concepts** of the Blockchain technology
- Give a **simple implementation** with the main functionalites of a Blockchain
- **Easy to configure and run** in your computer
- Self documented and **easy to hack/extend**

## Bitcoin in a nutshell
Bitcoin is decentralized network implementing a peer-to-peer electronic cash system. Bitcoin uses blockchain to keep track of all the account balances of the users in the network. 

**Identities**: Bitcoin is based on public key cryptography. Each user has his/her own public and private key that are used to identify them and validate transactions.

**Transactions**: Transactions are the way of transfering value from one user to other. A transaction have the informatio of who is giving the value and to whom. In order to be sure that a transaction is valid we need to check two things:
- *Who is transfering the value actually has that value*: That means that he recived that value from a previous transaction and he did not spent it yet.
- *Who is transfering the value is who created the transcation*: That means that he signed the transaccion. It can be easly verified with his public key.

**Blocks**: Blocks are basically a bunch of valid transactions. The nodes of the network receive which are the pending transactions and validate them. Then they start working in the proof-of-work. To make it hard to accept blocks a valid block has to have a *hash* with a concrete properties. Those properties are usually "start with *k* zeros". But the hash of something remains always equal... so what we do is add some extra data to the block (the nonce) and change it until we get a block which a hash value meeting the property that we were looking for. As soon as we find it we broadcast the block and the other nodes will put it in the blockchain and will keep working from there.

**Incentives**: A transaction is accepted when it is finally in the blockchain, so it is an a valid block. As we saw validating blocks is quite expensive, due to the proof-of-work procedure. So, why should one join the bitcoin network and validate blocks? The answer is that the user who finds the next valid block for the blockchain creates a new transaccion from which he get some bitcoins.

## Digital Arquitecture

The arquitecture of this project is heavly based on the one specified in the [Bitcoin whitepaper](https://bitcoin.org/bitcoin.pdf). Because of that this project can be used to grasp the concepts behind Bitcoin easier.

### Identities
As we have seen identities are pairs of private and public keys. In Digitial the logic for the managment of identities is in the [crypto.go](https://github.com/jomsdev/digital/blob/master/core/crypto.go) file. Digital uses eliptic curves for the keys, the P256 for being more exact. In Golang cryptographic operations for P256 are implemented using constant-time algorithms.

An identity is represented as a Keypair

```go
type Keypair struct {
	Public  []byte
	Private []byte
}
```

To transform Keys into something "humar readable" digital uses Base64 for byte encoding.

### Transactions

```go
type Transaction struct {
	Header    TransactionHeader
	Signature []byte
	Payload   []byte
}
```

```go
type TransactionType uint32

const (
	Payment TransactionType = 0
	CoinCreation
)
```

```go
type TransactionHeader struct {
	From        []byte
	To          []byte
	Type        TransactionType
	Timestamp   uint32
	PayloadHash []byte
}
```

### Blocks

```go
type Block struct {
	Header       BlockHeader
	Signature    []byte
	Transactions []Transaction
}
```

```go
type BlockHeader struct {
	From       []byte
	PrevHash   []byte
	MerkleRoot []byte
	Timestamp  uint32
	Nonce      uint32
}
```
### Blockchain

```go
type Blockchain struct {
	Identities []string
	Blocks     []Block
}
```

## Example


```go
// TODO: Example of Blockchain + three identities + two transaction + one block
```

## License

Digital is under the MIT license. See the [LICENSE](https://github.com/jomsdev/digital/blob/master/LICENSE) file for details.
