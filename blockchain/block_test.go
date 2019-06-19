package blockchain

import (
	"crypto/dsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/mourtisma/go-blockchain-api/blockchain/keyStore"
	"testing"
)

func TestComputeBlockHash(t *testing.T) {
	previousBlock := Block{
		"prevBlock",
		"abcd",
		nil,
		nil,
		nil,
	}

	block := Block{
		"block",
		"",
		&previousBlock,
		nil,
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

func TestComputeDigitalSignature(t *testing.T) {
	keyStore := keyStore.KeyStore{}
	keyStore.GenerateKeys()

	block := Block{
		"block",
		"abcd",
		nil,
		nil,
		nil,
	}

	hash := []byte(block.blockHash)

	signature, _ := block.ComputeDigitalSignature(keyStore)
	verifystatus := dsa.Verify(keyStore.PublicKey, hash, signature.r, signature.s)

	if verifystatus != true {
		t.Errorf("Block has failed verification while it shouldn't")
	}

}

func TestLinkToPreviousBlock(t *testing.T) {
	previousBlock := Block{
		"prevBlock",
		"abcd",
		nil,
		nil,
		nil,
	}

	block := Block{
		"block",
		"",
		&previousBlock,
		nil,
		nil,
	}

	block.LinkToPreviousBlock()

	if previousBlock.nextBlock != &block {
		t.Errorf("Blocks were not linked correctly")
	}

}
