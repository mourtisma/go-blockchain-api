package blockchain


import (
	"testing"
	"crypto/sha256"
	"encoding/base64"
)

func TestCreateGenesisBlock(t *testing.T) {
	data := "hello world"
	h := sha256.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	expectedHash := base64.StdEncoding.EncodeToString(hash)
	
	blockChain := BlockChain{}
	block := blockChain.CreateGenesisBlock("hello world")

	if block.nextBlock != nil {
		t.Errorf("Genesis block shouldn't point to a block when created")	
	}

	if block.previousBlock != nil {
		t.Errorf("Genesis block shouldn't be pointed by a block")	
	}

	if actualHash := block.blockHash; actualHash != expectedHash {
		t.Errorf("Genesis block got wrong hash: got %v want %v",
		actualHash, expectedHash)
	}

}

func TestSaveBlock(t *testing.T) {
	firstBlock := Block{
		"firstBlock",
		"abcd",
		nil,
		nil,
	}

	h := sha256.New()
	h.Write([]byte("abcdsecondBlock"))
	hash := h.Sum(nil)
	expectedHash := base64.StdEncoding.EncodeToString(hash)

	blockChain := BlockChain{
		[]*Block{&firstBlock},
	}

	newBlock := Block{
		"secondBlock",
		"",
		&firstBlock,
		nil,
	}

	blockChain.SaveBlock(&newBlock)

	// Validate hash
	if actualHash := newBlock.blockHash; actualHash != expectedHash {
		t.Errorf("New block got wrong hash: got %v want %v",
		actualHash, expectedHash)
	}

	// Validate chain size
	if len(blockChain.blocks) != 2 {
		t.Errorf("Wrong chain size: got %v want %v",
		len(blockChain.blocks), 2)
	}

	// Validate second block adress
	secondBlock := blockChain.blocks[1]
	if secondBlock != &newBlock {
		t.Errorf("Wrong second block")
	}

	// Validate links between blocks
	if secondBlock.previousBlock != &firstBlock {
		t.Errorf("First block is not pointed by second block")
	}
	if firstBlock.nextBlock != secondBlock {
		t.Errorf("Second block is not pointed by first block")
	}


}

func TestIsValid(t *testing.T) {

	firstBlock := Block{
		"firstBlock",
		"",
		nil,
		nil,
	}

	h := sha256.New()
	
	h.Write([]byte("firstBlock"))
	firstHash := h.Sum(nil)
	firstBlockHash := base64.StdEncoding.EncodeToString(firstHash)
	firstBlock.blockHash = firstBlockHash

	secondBlock := Block{
		"secondBlock",
		"",
		&firstBlock,
		nil,
	}
	
	h.Reset()
	h.Write([]byte(firstBlockHash+"secondBlock"))
	secondHash := h.Sum(nil)
	secondBlockHash := base64.StdEncoding.EncodeToString(secondHash)
	secondBlock.blockHash = secondBlockHash

	firstBlock.nextBlock = &secondBlock

	blockChain := BlockChain{
		[]*Block{&firstBlock, &secondBlock},
	}

	if !blockChain.IsValid(&firstBlock) {
		t.Errorf("Chain should be valid but is not")
	}
}

func TestIsCorrupt(t *testing.T) {

	firstBlock := Block{
		"firstBlock",
		"abcd",
		nil,
		nil,
	}

	h := sha256.New()
	
	h.Write([]byte("firstBlock"))
	firstHash := h.Sum(nil)
	firstBlockHash := base64.StdEncoding.EncodeToString(firstHash)
	firstBlock.blockHash = firstBlockHash

	secondBlock := Block{
		"secondBlock",
		"",
		&firstBlock,
		nil,
	}
	
	secondBlock.blockHash = "deadbeef"

	firstBlock.nextBlock = &secondBlock

	blockChain := BlockChain{
		[]*Block{&firstBlock, &secondBlock},
	}

	if blockChain.IsValid(&firstBlock) {
		t.Errorf("Chain shouldn't be valid but it is")
	}
}