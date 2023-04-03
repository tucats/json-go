package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func main() {
	var err error

	text := ""
	typeName := ""
	input := os.Stdin
	camel := false
	outfile := ""

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "-h", "--help":
			help()

		case "-c", "--camel-case", "--camel":
			camel = true

		case "-f", "--file":
			name := os.Args[i+1]
			input, err = os.Open(name)
			i++

		case "-o", "--output":
			outfile = os.Args[i+1]
			i++

		case "-t", "--type":
			typeName = os.Args[i+1]
			i++

		default:
			err = fmt.Errorf("unrecognized option: %s", arg)
		}

		if err != nil {
			break
		}
	}

	b := []byte{}

	if err == nil {
		for {
			d := make([]byte, 16384)
			count := 0

			count, err = input.Read(d)
			if count == 0 {
				if err == io.EOF {
					err = nil
				}

				break
			}

			b = append(b, d[:count]...)
		}
	}

	if err == nil {
		text, err = Convert(b, typeName, camel)
	}

	if err == nil {
		if outfile == "" {
			fmt.Println(text)
		} else {
			err = os.WriteFile(outfile, []byte(text), fs.ModePerm)
		}
	}

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
