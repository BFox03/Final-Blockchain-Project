package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)


/*
	This file contains the functions necessary to make and continue the blockchain

	A blockchain is a connection of blocks through a validation of hashes
	This is considered to be the future of data storage because blockchains do not require a centralized server
	Blockchains may be stored on third-party nodes allowing for a community to own the data rather than one manager

	Each function is defined and serves a purpose to each other
	These functions include:
	Hashing (unique value that only applies to ,
	Making Block,
	Appending Blocks,
	and Creating a Genesis Block
 */


// Block Declare what makes up a block: the time it was hashed (integer 64 byte type), data of the block (byte slice type), the previous hash, the current hash, and the nonce (integer)
type Block struct {
	Time int64
	Data []byte
	PrevHash []byte
	Hash []byte
	Nonce int
}

// Blockchain Declare what makes up the blockchain. An array of blocks
type Blockchain struct {
	blocks []*Block
}

// MakeHash This function grabs all the info (previous hash, hash and time) combines them into a byte slice array and hashes it
func (b *Block) MakeHash() {
	time := []byte(strconv.FormatInt(b.Time, 10))
	toHash := bytes.Join([][]byte{
		b.PrevHash,
		b.Hash,
		time},
		[]byte{})
	hash := sha256.Sum256(toHash)

	b.Hash = hash[:]
}

//MakeBlock Takes all the data given to the block and puts it into the block struct type. As well this function calls the proof of work function to validate it
func MakeBlock(data string, prevHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevHash,
		[]byte{},
		0}
	pow := NewPoW(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//AddBlock takes the data given in the main sends it to MakeBlock and appends the new block onto the blockchain array
func (bc *Blockchain) AddBlock(data string){
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := MakeBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//GenesisBlock All blockchains begin with a genesis block which is the first in the chain
func GenesisBlock() *Block {
	return MakeBlock("Genesis Block", []byte{})
}

//MakeBlockchain initializes the datatype creating the array of blocks where information will be stored
func MakeBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}