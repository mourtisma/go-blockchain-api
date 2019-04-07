package blockchain

import (
	"crypto/sha256"
	"encoding/base64"
)

// IBlock is the generic interface implemented by the blocks
type IBlock interface {
	ComputeBlockHash() string
	SetBlockHash(blockHash string)
	LinkToPreviousBlock()
}

// Block struct represents a chain's block
type Block struct {
	data          string
	blockHash     string
	previousBlock *Block
	nextBlock     *Block
}

// ComputeBlockHash calculates the base-64 encoding of the SHA-256 block checksum
func (block *Block) ComputeBlockHash() string {
	h := sha256.New()
	if (block.previousBlock == nil) {
		h.Write([]byte(block.data))
	} else {
		previousBlock := block.previousBlock
		h.Write([]byte(previousBlock.blockHash + block.data))
	}

	hash := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash)
}

// SetBlockHash is a simple setter
func (block *Block) SetBlockHash(blockHash string) {
	block.blockHash = blockHash
}

// LinkToPreviousBlock takes the block's predecessor,
// and links it to the current block
func (block *Block) LinkToPreviousBlock() {
	if block.previousBlock != nil {
		previousBlock := block.previousBlock
		previousBlock.nextBlock = block
	}
}
