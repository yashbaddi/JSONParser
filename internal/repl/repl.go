package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yashbaddi/jsonparser/internal/json"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(out, err)
	}

	for {
		fmt.Fprint(out, "JSON Parser > ")

		if !scanner.Scan() {
			fmt.Fprintln(out, "Scan failed")
			return
		}

		input := scanner.Bytes()
		data, ok := json.JSONParser(input)

		if !ok {
			fmt.Fprintln(out, "Error Processing the JSON ")
		}

		fmt.Fprintf(out, "JSON Output: \n %v\n", data)

		fmt.Fprintf(out, "JSON Parsed Successfully!! \n ")

	}
}
