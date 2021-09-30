package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

var outputs []string

func init() {
	if len(os.Args) == 1 {
		printHelp()
	}
}

func printHelp() {
	fmt.Println(help)
	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	for i, val := range args {
		switch strings.ToLower(val) {
		case "base64e":
			input := args[i+1]
			outputs = append(outputs, fmt.Sprintf("base64e:\n%v\n\noutput:\n%v\n", input, base64.StdEncoding.EncodeToString([]byte(input))))
		case "base64d":
			input := args[i+1]
			decodeString, err := base64.StdEncoding.DecodeString(input)
			if err != nil {
				outputs = append(outputs, fmt.Sprintf("base64d:\n%v\n\noutput:\n%v\n", input, "FAIL: "+err.Error()))
			} else {
				outputs = append(outputs, fmt.Sprintf("base64d:\n%v\n\noutput:\n%v\n", input, string(decodeString)))
			}
		}
	}

	for _, val := range outputs {
		fmt.Println("---------------------------------------------------------------------------------------------")
		fmt.Print(val)
	}
}
