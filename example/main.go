package main

import (
	"github.com/tkc/rootdir"
	"log"
)

func main() {
	root := rootdir.Root()
	log.Print(root)
}
