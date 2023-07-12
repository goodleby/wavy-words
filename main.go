package main

import (
	"log"
	"os"

	"github.com/goodleby/wavy-words/wavyword"
)

func main() {
	wavyWord, err := wavyword.New("reallylongwavyword", 3)
	if err != nil {
		log.Printf("error creating new wavy word: %v", err)
		os.Exit(1)
	}

	wavyWord.Print()
}
