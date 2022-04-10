package main

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func Genesis() *Block {
	return CreateBlock("genesis", []byte{})
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (b *Block) DeriveHash() {
	// thi s will join our previous block 's  relevant info with
	// new blocks
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// this performs the actual hashin algorithm
	hash := sha256.Sum256(info)
	//  referencing the storage of on  b.Hash
	b.Hash = hash[:]
}

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}
