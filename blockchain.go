package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents each 'block' in the blockchain
type Block struct {
	Index        int
	Timestamp    time.Time
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
}

// CalculateHash calculates the SHA256 hash of the block
func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp.String() + b.Data + b.PrevHash + string(b.Nonce)
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

// Blockchain is a series of validated blocks
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain creates a new Blockchain with the genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := &Block{Index: 0, Timestamp: time.Now(), Data: "Genesis Block", PrevHash: "", Hash: ""}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return &Blockchain{blocks: []*Block{genesisBlock}}
}

// GetLastBlock retrieves the most recent block in the chain
func (bc *Blockchain) GetLastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.GetLastBlock()
	newBlock := &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.MineBlock(2) // difficulty level
	bc.blocks = append(bc.blocks, newBlock)
}

// MineBlock performs proof-of-work
func (b *Block) MineBlock(difficulty int) {
	prefix := bytes.Repeat([]byte("0"), difficulty)
	for {
		b.Hash = b.CalculateHash()
		if bytes.HasPrefix([]byte(b.Hash), prefix) {
			fmt.Printf("Block mined with hash: %s\n", b.Hash)
			break
		}
		b.Nonce++
	}
}

// IsValid checks if the blockchain is valid
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		prevBlock := bc.blocks[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")

	for _, block := range bc.blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Println()
	}

	fmt.Printf("Blockchain valid? %v\n", bc.IsValid())
}
