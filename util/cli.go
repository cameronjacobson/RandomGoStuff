// Copyright © 2013 Cameron Jacobson

// Description:
//  Simple skeleton file for creating cli pipe utilities

// Usage:  
//   $ cat input.txt | go run cli.go > output.txt
//   $ cat input.txt | compiled > output.txt

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			fmt.Println("ERROR:",err);
		}

		// DO STUFF

		if _, err = writer.WriteString(line); err != nil {
			fmt.Println("ERROR:",err)
		}
	}
	writer.Flush()
}
