package blockchain

import (
	"github.com/mourtisma/go-blockchain-api/blockchain/keyStore"
	"crypto/dsa"
)

// IBlockChain is the generic interface implemented by our BlockChain model
type IBlockChain interface {
	SaveBlock(bock Block)
}

// BlockChain struct represent a basic BlockChain model
type BlockChain struct {
	keyStore keyStore.KeyStore
	blocks []*Block
}

// CreateGenesisBlock starts the chain by creating its header block
func (blockChain *BlockChain) CreateGenesisBlock(data string) Block {
	block := Block{data, "", nil, nil, nil,}
	blockHash := block.ComputeBlockHash()
	block.SetBlockHash(blockHash)
	block.Sign(blockChain.keyStore)
	return block
}

// SaveBlock calculates a block hash given its previous block,
// creates the link between the two
// and inserts it in the chain
func (blockChain *BlockChain) SaveBlock(block *Block) {
	if len(block.blockHash) <= 0 {
		blockHash := block.ComputeBlockHash()
		block.SetBlockHash(blockHash)
	}

	if (block.digitalSignature == nil) {
		block.Sign(blockChain.keyStore)
	}

	block.LinkToPreviousBlock()
	blockChain.blocks = append(blockChain.blocks, block)
}

// IsValid checks, given a block in the chain,
// its expected hash against its current hash
// then carries on the check for its successors to assert
// the chain's integrity
func (blockChain *BlockChain) IsValid(block *Block) bool {
	blockHash := block.ComputeBlockHash()
	signature, _ := block.ComputeDigitalSignature(blockChain.keyStore)

	if block.nextBlock == nil {
		return blockHash == block.blockHash && dsa.Verify(blockChain.keyStore.PublicKey, 
														  []byte(blockHash), 
														  signature.r, 
														  signature.s)
	}

	hashAndSignatureValid := blockHash == block.blockHash && dsa.Verify(blockChain.keyStore.PublicKey, 
		[]byte(blockHash), 
		signature.r, 
		signature.s)
	return hashAndSignatureValid && blockChain.IsValid(block.nextBlock)
}
