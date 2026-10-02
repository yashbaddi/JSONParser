package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yashbaddi/jsonparser/internal/json"
	"github.com/yashbaddi/jsonparser/internal/repl"
)

func isFileFlagSet() bool {

	fileFlagSet := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "f" || f.Name == "file" {
			fileFlagSet = true
		}
	})

	return fileFlagSet

}

func fileBuffer(filePath string) ([]byte, error) {

	return os.ReadFile(filePath)

}

func main() {
	var filePath string

	flag.StringVar(&filePath, "f", "", "path to the file (shorthand)")
	flag.StringVar(&filePath, "file", "", "path to the file")
	flag.Parse()

	if filePath != "" {
		filebuf, err := fileBuffer(filePath)
		if err != nil {
			fmt.Println("CLI Error:", err)
			os.Exit(2)
		}

		parsedJSON, ok := json.JSONParser(filebuf)
		if !ok {
			fmt.Println("Error Parsing JSON File")
			os.Exit(1)
		}

		fmt.Println("Parsed Success")
		fmt.Print(parsedJSON)
		return
	}

	repl.Start(os.Stdin, os.Stdout)
}
