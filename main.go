package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFilename string = `shakespeare-only.txt`

var trie *Trie

func init() {
	f, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	trie = NewTrie()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := Parse(line)
		for _, token := range tokens {
			trie.Insert(token)
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("imported file %q, parsed %d words", inputFilename, len(trie.Words))
}

func main() {
	fmt.Println("ohai")
}
