package main

import (
	"fmt"

	"github.com/againagainst/conv/pkg/conv"
)

func main() {
	input := conv.ParseArgs()
	output, err := conv.LoadPlugin(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output.Sprint())
	}
}
