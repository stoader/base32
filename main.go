package main

import (
	"encoding/base32"
	"fmt"
	"os"
	"flag"
	"io/ioutil"
)



func encode(input []byte) ([]byte) {
	encoding := base32.StdEncoding

	inputLen := len(input)
	outputLen := encoding.EncodedLen(inputLen)

	output := make([]byte, outputLen)

	encoding.Encode(output, input)

	return output
}

func decode(input []byte) ([]byte) {
	encoding := base32.StdEncoding

	inputLen := len(input)
	outputLen := encoding.DecodedLen(inputLen)

	output := make([]byte, outputLen)

	_, err := encoding.Decode(output, input)
	if err != nil {
		panic(err)
	}


	return output
}

func main() {
	var isDecode bool
	flag.BoolVar(&isDecode, "d",false, "base32 decode")

	flag.Parse()

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var output []byte
	if isDecode {
		output = decode(input)
	} else {
		output = encode(input)
	}

	fmt.Print(string(output))
}

