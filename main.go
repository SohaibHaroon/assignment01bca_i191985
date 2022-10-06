package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

var HashofLastBlock string = ""

// a structure denoting a Block of BlockChain. It contains Transaction, Hash of Block, Previous Block Hash & Nonce.
type block struct {
	transaction       string
	hashOfBlock       string
	previousBlockHash string
	nonce             int
}

// a structure of pointer to an array containing Blocks of Block Chain
type BlockChainList struct {
	bc_list []*block
}

// function to create a new block. The block hash is then computed and block is returned after being added in array.
func (b *BlockChainList) NewBlock(b_transaction string, b_nonce int, b_previousHash string) *block {

	newBlock := new(block)
	newBlock.transaction = b_transaction
	newBlock.nonce = b_nonce
	newBlock.previousBlockHash = b_previousHash
	hashstring := strconv.Itoa(newBlock.nonce) + newBlock.previousBlockHash + newBlock.transaction
	newBlock.hashOfBlock = CalculateHash(hashstring)
	b.bc_list = append(b.bc_list, newBlock)

	lengthOfBlockArray := len(b.bc_list)
	lengthOfBlockArray = lengthOfBlockArray - 1
	hashstring = strconv.Itoa(b.bc_list[lengthOfBlockArray].nonce) + b.bc_list[lengthOfBlockArray].previousBlockHash + b.bc_list[lengthOfBlockArray].transaction
	HashofLastBlock = CalculateHash(hashstring)

	return newBlock
}

// Change Block Function, takes the block number to be changed and amends the transaction. After this new hash of Block is calculated.
func ChangeBlock(b *BlockChainList, index int, transaction_b string) *block {
	index = index - 1
	b.bc_list[index].transaction = transaction_b
	BlockHash := strconv.Itoa(b.bc_list[index].nonce) + b.bc_list[index].previousBlockHash + b.bc_list[index].transaction
	b.bc_list[index].hashOfBlock = CalculateHash(BlockHash)
	modifiedBlock := b.bc_list[index]
	return modifiedBlock
}

// This function verifies Block Chain by comparing the Hash of Block and Previous Block Hash. If changed then notified.
func VerifyChain(b *BlockChainList) int {
	var i = 0
	var flag = 0
	for i = 0; i < (len(b.bc_list) - 1); i++ {
		if b.bc_list[i].hashOfBlock != b.bc_list[i+1].previousBlockHash {
			flag = 1
			fmt.Println("Changes detected in Block ", i+1)
		}
	}

	// computing the final block Hash and comparing it with the saved one already.
	lenOfBlockChain := len(b.bc_list)
	lenOfBlockChain = lenOfBlockChain - 1
	hashstring := strconv.Itoa(b.bc_list[lenOfBlockChain].nonce) + b.bc_list[lenOfBlockChain].previousBlockHash + b.bc_list[lenOfBlockChain].transaction
	LatestHashofLastBlock := CalculateHash(hashstring)

	if LatestHashofLastBlock != HashofLastBlock {
		fmt.Println("Changes detected in Block ", lenOfBlockChain+1)
	}

	return flag
}

// Function to calculate sha256 Hash of Block
func CalculateHash(stringToHash string) string {
	blockHash := sha256.Sum256([]byte(stringToHash))
	calculatedHash := fmt.Sprintf("%x", blockHash)
	return calculatedHash
}

// This function takes the array of Blocks and print the Blocks in a decent manner
func ListBlocks(blocks *BlockChainList) {
	for i := 0; i < len(blocks.bc_list); i++ {
		fmt.Println("--------------------------- Block ", i+1, "---------------------------")
		fmt.Println()
		fmt.Println("Transaction:     ", blocks.bc_list[i].transaction)
		fmt.Println("Nonce:           ", blocks.bc_list[i].nonce)
		fmt.Println("Block Hash:      ", blocks.bc_list[i].hashOfBlock)
		fmt.Println("Prev Block Hash: ", blocks.bc_list[i].previousBlockHash)
		fmt.Println()
	}
}

func main() {
	listOfBlock := new(BlockChainList)
	// Block#01
	listOfBlock.NewBlock("Alice to Bob", 10, "")
	prevBlockHash := strconv.Itoa(listOfBlock.bc_list[0].nonce) + listOfBlock.bc_list[0].previousBlockHash + listOfBlock.bc_list[0].transaction
	calculatePrevBlockHash := CalculateHash(prevBlockHash)
	// Block#02
	listOfBlock.NewBlock("Alice to Charlie", 20, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[1].nonce) + listOfBlock.bc_list[1].previousBlockHash + listOfBlock.bc_list[1].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)
	// Block#03
	listOfBlock.NewBlock("Alice to Dany", 30, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[2].nonce) + listOfBlock.bc_list[2].previousBlockHash + listOfBlock.bc_list[2].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)
	// Block#04
	listOfBlock.NewBlock("Bob to Alice", 17, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[3].nonce) + listOfBlock.bc_list[3].previousBlockHash + listOfBlock.bc_list[3].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)
	// Block#05
	listOfBlock.NewBlock("Bob to Dany", 27, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[4].nonce) + listOfBlock.bc_list[4].previousBlockHash + listOfBlock.bc_list[4].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)
	// Block#06
	listOfBlock.NewBlock("Bob to Malory", 37, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[5].nonce) + listOfBlock.bc_list[5].previousBlockHash + listOfBlock.bc_list[5].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)
	// Block#07
	listOfBlock.NewBlock("Charlie to Alice", 11, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[6].nonce) + listOfBlock.bc_list[6].previousBlockHash + listOfBlock.bc_list[6].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)
	// Block#08
	listOfBlock.NewBlock("Charlie to Bob", 21, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[7].nonce) + listOfBlock.bc_list[7].previousBlockHash + listOfBlock.bc_list[7].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)
	// Block#09
	listOfBlock.NewBlock("Charlie to Dany", 31, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[8].nonce) + listOfBlock.bc_list[8].previousBlockHash + listOfBlock.bc_list[8].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)
	// Block#10
	listOfBlock.NewBlock("Dany to Malory", 24, calculatePrevBlockHash)
	prevBlockHash = strconv.Itoa(listOfBlock.bc_list[9].nonce) + listOfBlock.bc_list[9].previousBlockHash + listOfBlock.bc_list[9].transaction
	calculatePrevBlockHash = CalculateHash(prevBlockHash)

	// PRINTING BLOCKS WITHOUT CHANGING ANYONE
	ListBlocks(listOfBlock)

	// Change Block Function
	// Changing transaction of Block No 2.
	var block_no = 2
	newModifiedBlock := ChangeBlock(listOfBlock, block_no, "Alice to Malory")
	fmt.Println()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("                           Changing the Block                         ")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()

	fmt.Println("The New Modified Block is: ")
	fmt.Println()
	fmt.Println("Block Number:        ", block_no)
	fmt.Println("Transaction:         ", newModifiedBlock.transaction)
	fmt.Println("Nonce:               ", newModifiedBlock.nonce)
	fmt.Println("Block Hash:          ", newModifiedBlock.hashOfBlock)
	fmt.Println("Previous Block Hash: ", newModifiedBlock.previousBlockHash)
	fmt.Println()

	// Changing transaction of Block No 8.
	block_no = 8
	newModifiedBlock = ChangeBlock(listOfBlock, block_no, "Charlie to Malory")
	fmt.Println()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("                           Changing the Block                         ")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()

	fmt.Println("The New Modified Block is: ")
	fmt.Println()
	fmt.Println("Block Number:        ", block_no)
	fmt.Println("Transaction:         ", newModifiedBlock.transaction)
	fmt.Println("Nonce:               ", newModifiedBlock.nonce)
	fmt.Println("Block Hash:          ", newModifiedBlock.hashOfBlock)
	fmt.Println("Previous Block Hash: ", newModifiedBlock.previousBlockHash)
	fmt.Println()

	// Changing transaction of Block No 10 (Final Block).
	block_no = 10
	newModifiedBlock = ChangeBlock(listOfBlock, block_no, "Dany to Simba")
	fmt.Println()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("                           Changing the Block                         ")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()

	fmt.Println("The New Modified Block is: ")
	fmt.Println()
	fmt.Println("Block Number:        ", block_no)
	fmt.Println("Transaction:         ", newModifiedBlock.transaction)
	fmt.Println("Nonce:               ", newModifiedBlock.nonce)
	fmt.Println("Block Hash:          ", newModifiedBlock.hashOfBlock)
	fmt.Println("Previous Block Hash: ", newModifiedBlock.previousBlockHash)
	fmt.Println()

	// Verifying Block Function
	fmt.Println()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("                           Verifying the BlockChain                         ")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()
	change_detected := VerifyChain(listOfBlock)
	if change_detected == 0 {
		fmt.Println("BlockChain Verified! No changes detected...")
	}

}
