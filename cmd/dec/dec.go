package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/sh3rp/encdec"
)

var dec string

func main() {
	flag.StringVar(&dec, "d", "b64", "Encoding to use [b64|b32|url|hex]")
	flag.Parse()

	var decoded []byte
	var err error

	if flag.Arg(0) == "" {
		reader := bufio.NewReader(os.Stdin)
		data := make([]byte, 2<<20)
		numRead, err := reader.Read(data)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		decoded, err = encdec.Codecs[dec].Decode(data[:numRead])
	} else {
		decoded, err = encdec.Codecs[dec].Decode([]byte(flag.Arg(0)))
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%s", string(decoded))
	}
}
