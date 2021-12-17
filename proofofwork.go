package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)


/*
	This file contains the functions and values necessary for a proof of work protocol

	A proof of work function is necessary in most blockchains to avoid an overflow of blocks
	By slowing down the process at which blocks can be created, taking over a blockchain by changing blocks becomes more difficult

	The functions below do this process by requiring the hash of the block to reach a certain amount of leading zeros
	Doing this forces the hashing function to continue over and over until it reaches the difficulty set
 */


//Used to avoid overflow error
var maxNonce = math.MaxInt

/*
	Difficulty of mining the block (smaller number will decrease the time necessary to mine)
	This does this by changing the amount of leading zeros needed for the hash
	Currently, a minimum of 2 zeros is necessary (will produce a run time of around 15 seconds until it calculates the hashes, depends on computer speed)
 */
const targetSize = 10

//ProofOfWork This struct type contains the pointer to the block and a pointer to the difficulty necessary (we will compare to calculated hashes)
type ProofOfWork struct {
	block *Block
	target *big.Int
}

//NewPoW This function creates the target value that must be hit in order for the block to be valid
func NewPoW(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetSize))

	pow := &ProofOfWork{b, target}

	return pow
}

/*
	IntToHex Will convert the nonce into a value that can be hashed
	The function does this, so it may have another value to hash until it reaches the specified difficulty
 */
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

//SetData Creates the data to send over to get hashed
func (pow *ProofOfWork) SetData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
		pow.block.PrevHash,
		pow.block.Data,
		IntToHex(pow.block.Time),
		IntToHex(int64(targetSize)),
		IntToHex(int64(nonce))},
		[]byte{})

	return data
}

/*
	Run is the meat of the proof of work file
	The function prepares the data, hashes it, converts the hash into a big int type, and compares it wil the target
	If the hash does not meet the requirements of the difficulty, it will up the nonce (creating a new hash to check) and loop
 */
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInteger big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining block /%s/\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.SetData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("/r%x\n", hash)
		hashInteger.SetBytes(hash[:])

		if hashInteger.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println("All work done! :D")

	return nonce, hash[:]
}

/*
	Valid is a very simple function that checks whether the difficulty has been reached
	The return of this statement is what is displayed in the terminal for whether it is valid or not
 */
func (pow *ProofOfWork) Valid() bool {
	var hashInteger big.Int

	data := pow.SetData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInteger.SetBytes(hash[:])

	validity := hashInteger.Cmp(pow.target) == -1

	return  validity
}