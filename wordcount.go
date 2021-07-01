package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
        fmt.Println("usage: wordcount <sentence>")
        os.Exit(1)
    }
	s := strings.Trim(os.Args[1], "\"'")
	words := strings.Fields(s)
	fmt.Println(len(words))
}
