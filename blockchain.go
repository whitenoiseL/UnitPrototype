package main

import "./github.com/syndtr/goleveldb/leveldb"

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := leveldb.OpenFile("./db", nil)



	return &Blockchain{[]*Block{NewGensisBlock()}}
}