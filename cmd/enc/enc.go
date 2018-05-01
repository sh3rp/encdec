package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/sh3rp/encdec"
)

var enc string

func main() {
	flag.StringVar(&enc, "s", "b64", "Encoding to use [b64|b32|url]")
	flag.Parse()

	var encoded []byte
	var err error

	if flag.Arg(0) == "" {
		reader := bufio.NewReader(os.Stdin)
		data := make([]byte, 2^20)
		numRead, err := reader.Read(data)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		encoded, err = encdec.Codecs[enc].Encode(data[:numRead])
	} else {
		encoded, err = encdec.Codecs[enc].Encode([]byte(flag.Arg(0)))
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%s", string(encoded))
	}
}
