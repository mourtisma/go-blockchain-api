package blockchain

import (
	"errors"
	"crypto/sha256"
	"crypto/rand"
	"crypto/dsa"
	"encoding/base64"
	"blockchain/keyStore"
	"math/big"
)

// IBlock is the generic interface implemented by the blocks
type IBlock interface {
	ComputeBlockHash() string
	SetBlockHash(blockHash string)
	LinkToPreviousBlock()
	ComputeDigitalSignature()
	Sign()
}

// Block struct represents a chain's block
type Block struct {
	data          string
	blockHash     string
	previousBlock *Block
	nextBlock     *Block
	digitalSignature *DigitalSignature
}

type DigitalSignature struct {
	r, s *big.Int
}

// ComputeBlockHash calculates the base-64 encoding of the SHA-256 block checksum
func (block *Block) ComputeBlockHash() string {
	h := sha256.New()
	if block.previousBlock == nil {
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

// ComputeDigitalSignature computes the digital signature of a block, given a KeyStore
func (block *Block) ComputeDigitalSignature(ks keyStore.KeyStore) (*DigitalSignature, error) {
	r, s, e := dsa.Sign(rand.Reader, ks.PrivateKey, []byte(block.blockHash))
	if e != nil {
		err := errors.New("Error signing the hash")
		return nil, err
	}
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	return &DigitalSignature{r, s}, nil 
}

// Sign computes the digital signature of a block, given a KeyStore, and signs it
func (block *Block) Sign(ks keyStore.KeyStore) {
	signature, _ := block.ComputeDigitalSignature(ks)
	block.digitalSignature = signature
}

// LinkToPreviousBlock takes the block's predecessor,
// and links it to the current block
func (block *Block) LinkToPreviousBlock() {
	if block.previousBlock != nil {
		previousBlock := block.previousBlock
		previousBlock.nextBlock = block
	}
}
