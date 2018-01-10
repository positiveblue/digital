# Digital - A basic blockchain


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

**Identities**: Bitcoin is based on public cryptography. Each user has his/her own public and private key that are used to identify them and validate transactions.

**Transactions**: Transactions are the way of transfering value from one user to other. A transaction have the informatio of who is giving the value and to whom. In order to be sure that a transaction is valid we need to check two things:
- Who is transfering the value actually has that value
- Who is transfering the value is who created the transcation

**Blocks**:



## Digital Arquitecture


## The future of blockchain

## License

Digital is under the MIT license. See the [LICENSE](https://github.com/jomsdev/digital/blob/master/LICENSE) file for details.
