package blockchain

import (
	"crypto/sha256"
	"encoding/base64"
	"testing"
)

func TestSetBlockHash(t *testing.T) {
	previousBlock := Block{
		"prevBlock",
		"abcd",
		nil,
		nil,
	}

	block := Block{
		"block",
		"",
		&previousBlock,
		nil,
	}

	h := sha256.New()
	h.Write([]byte("abcdblock"))
	hash := h.Sum(nil)
	expectedHash := base64.StdEncoding.EncodeToString(hash)
	actualHash := block.ComputeBlockHash()

	if actualHash != expectedHash {
		t.Errorf("Genesis block got wrong hash: got %v want %v",
			actualHash, expectedHash)
	}

}

func TestLinkToPreviousBlock(t *testing.T) {
	previousBlock := Block{
		"prevBlock",
		"abcd",
		nil,
		nil,
	}

	block := Block{
		"block",
		"",
		&previousBlock,
		nil,
	}

	block.LinkToPreviousBlock()

	if previousBlock.nextBlock != &block {
		t.Errorf("Blocks were not linked correctly")
	}

}
