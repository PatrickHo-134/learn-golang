package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G0 R0ck!"
	replacer := strings.NewReplacer("0", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
