package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

// this tool is used to capture the data sent to the stdin by the protoc and save it to a file for later use
func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the proto file from stdin: %s", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("./protorequest.bin", in, fs.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing the proto file from stdin: %s", err)
		os.Exit(1)
	}
}
