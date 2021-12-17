package main

import (
	"fmt"
	"strconv"
)


/*
	This code is a very very basic blockchain representation that can:
		Create and process blocks
		Check through a proof of work system to add mining time
		Change data input and see its output through console

	If you would like to change the inputs and difficulty:
		Inputs are at line 21 on main.go
		Difficulty variation is on line 16 proofofwork.go
 */


func main(){
	//Initialize the Blockchain, create one
	bc := MakeBlockchain()

	//Add data to the blocks
	bc.AddBlock("This is data that the Blockchain would store")
	bc.AddBlock("In this specific case, Data is stored as a string")
	bc.AddBlock("Each value given is that for the block, as you can see the previous hash is correct")
	bc.AddBlock("The data is hashed until it reaches a hash with 2 leading zeros, creating difficulty, implementing proof of work")

	/*
		Print out each block's prevHash, Data, and Hash
		Prints out all the blocks on the blockchain
	 */
	for _, block := range bc.blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		//The Proof of work function is done to check whether it meets the difficulty requirement
		pow := NewPoW(block)

		fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(pow.Valid()))
		fmt.Println("__________________")
	}
}