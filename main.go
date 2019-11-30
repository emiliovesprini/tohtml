package main

import (
	"fmt"
	"go/doc"
	"io/ioutil"
	"os"
)

func main() {
	var text string
	usage := "Usage: tohtml input.txt > output.html"
	switch len(os.Args) {
	default:
		// Catch non supported calls
		//
		fmt.Fprintln(os.Stderr, "Error: too many arguments.")
		fmt.Fprintln(os.Stderr, usage)
		return
	case 1:
		// Read stdin
		//
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin.")
			return
		}
		text = string(bytes)
	case 2:
		// Read file
		//
		bytes, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, "Error: file", os.Args[1], " does not exist.")
				fmt.Fprintln(os.Stderr, usage)
			} else {
				fmt.Fprintln(os.Stderr, "Error reading file.")
			}
			return
		}
		text = string(bytes)
	}
	// Call doc.ToHTML with an empty map, because there's no go code
	// from which to take identifiers to italicize. For details, see
	//
	//         go doc doc tohtml
	//
	doc.ToHTML(os.Stdout, text, make(map[string]string))
}
