package main

import (
	"fmt"
	"os"
	"autocorrectiontool"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Enter program name, input file, and output file")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error: Problem reading %s. \n", inputFile)
	}

	// Function calls will go here.
	data := autocorrectiontool.HexToDecimal(string(content))
	data = autocorrectiontool.BinToDecimal(data)
	data = autocorrectiontool.UpperCase(data)
	data = autocorrectiontool.LowerCase(data)
	data = autocorrectiontool.Capitalize(data)
	data = autocorrectiontool.ModifyText(data)
	data = autocorrectiontool.SetPunctuation(data)
	data = autocorrectiontool.ReplaceWithAn(data)


	err = os.WriteFile(outputFile, []byte(data), 0o644)
	if err != nil {
		fmt.Printf("Error: Trouble writing to %s. \n", outputFile)
	}
}
