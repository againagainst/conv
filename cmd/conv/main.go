package main

import (
	"fmt"
	"os"

	"github.com/againagainst/conv/pkg/conv"
)

func main() {
	arg := os.Args[1]
	result, err := conv.Ipv4ToHex(arg)
	if err != nil {
		fmt.Printf("%s: %s", err, result)
	} else {
		fmt.Println(result)
	}
}
