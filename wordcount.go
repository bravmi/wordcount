package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.Trim(os.Args[1], "\"'")
	words := strings.Fields(s)
	fmt.Println(len(words))
}
