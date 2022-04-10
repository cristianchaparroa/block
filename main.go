package main

import (
	bc "block/blockchain"
	"fmt"
)

func main() {
	chain := bc.NewBlockChain()
	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, b := range chain.GetBlocks() {
		fmt.Printf("Previous hash:  %x\n", b.PrevHash)
		fmt.Printf("data:  %s\n", b.Data)
		fmt.Printf("hash:   %x\n", b.Hash)
	}
}
