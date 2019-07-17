package main

import (
	"flag"
	"fmt"

	customFlag "github.com/abaehre/NumbersToWords/flag"
)

func main() {
	var intList customFlag.IntSlice
	flag.Var(&intList, "num", "Takes a list of integers to be outputted as english words, Usage: \"-num 1 -num 200\"")
	flag.Parse()
	fmt.Print(intList.String())
}
