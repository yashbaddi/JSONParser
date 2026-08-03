package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type JSONValue = any

var numberRE = regexp.MustCompile(`^-?([1-9](\d)*|0)(\.(\d)+)?([eE][+-]?(\d)+)?`)

func nullParser(input string) (any, string, bool) {
	rest := input
	if res := strings.HasPrefix(rest, "null"); res {
		return nil, rest[4:], true
	}
	return nil, input, false
}

func boolParser(input string) (bool, string, bool) {
	rest := input
	if res := strings.HasPrefix(rest, "true"); res {
		return true, rest[4:], true
	}
	if res := strings.HasPrefix(rest, "false"); res {
		return false, rest[5:], true
	}
	return false, input, false
}

func numberParser(input string) (float64, string, bool) {
	rest := input

	match := numberRE.FindString(rest)
	if match == "" {
		return 0, input, false
	}

	num, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, input, false
	}
	return num, rest[len(match):], true
}

func stringParser(input string) (string, string, bool) {
	rest := input

	if len(rest) == 0 || rest[0] != '"' {
		return "", input, false
	}

	var builder strings.Builder
	i := 1

	for i < len(rest) {

		switch rest[i] {
		case '"':
			return builder.String(), rest[i+1:], true

		case '\\':
			i++
			if i >= len(rest) {
				return "", input, false
			}
			r, ok := parseSplChars(input, &i)
			if !ok {
				return "", input, false
			}

			builder.WriteRune(r)

		default:
			builder.WriteByte(rest[i])

		}
		i++
	}

	return "", input, false

}

func parseSplChars(input string, i *int) (rune, bool) {
	rest := input

	switch rest[*i] {
	case '"':
		return '"', true
	case '\\':
		return '\\', true
	case '/':
		return '/', true
	case 'b':
		return '\b', true
	case 'f':
		return '\f', true
	case 'n':
		return '\n', true
	case 'r':
		return '\r', true
	case 't':
		return '\t', true
	case 'u':
		if *i+4 >= len(rest) {
			return rune(0), false
		}
		hex := rest[*i+1 : *i+5]
		value, err := strconv.ParseUint(hex, 16, 16)
		if err != nil {
			return rune(0), false
		}
		*i += 4

		return rune(value), true
	}

	return rune(0), false

}

func arrayParser(input string) ([]JSONValue, string, bool) {
	rest := input

	if !strings.HasPrefix(rest, "[") {
		return nil, input, false

	}

	var data []JSONValue
	rest = spaceParser(rest[1:])

	for !strings.HasPrefix(rest, "]") {
		rest = spaceParser(rest)
		var ok bool
		var value JSONValue
		value, rest, ok = valueParser(rest)
		if !ok {
			return nil, input, false
		}

		data = append(data, value)

		rest = spaceParser(rest)

		if strings.HasPrefix(rest, ",") {
			rest = rest[1:]
			continue
		}

		if strings.HasPrefix(rest, "]") {
			return data, rest[1:], true
		}

		return nil, input, false
	}
	return data, rest[1:], true

}

func objectParser(input string) (map[string]JSONValue, string, bool) {
	rest := input

	if !strings.HasPrefix(rest, "{") {
		return nil, input, false

	}

	data := make(map[string]any)
	rest = spaceParser(rest[1:])

	for !strings.HasPrefix(rest, "}") {
		rest = spaceParser(rest)

		var key string
		var ok bool
		key, rest, ok = stringParser(rest)
		if !ok {
			return nil, input, false
		}

		rest = spaceParser(rest)

		if !strings.HasPrefix(rest, ":") {
			return nil, input, false
		}

		var value JSONValue
		value, rest, ok = valueParser(rest[1:])
		if !ok {
			return nil, input, false
		}
		data[key] = value
		rest = spaceParser(rest)

		if strings.HasPrefix(rest, ",") {
			rest = rest[1:]
			continue
		}

		if strings.HasPrefix(rest, "}") {
			return data, rest[1:], true
		}

		return nil, input, false
	}
	return data, rest[1:], true
}

func spaceParser(input string) string {
	return strings.TrimLeftFunc(input, unicode.IsSpace)
}

func valueParser(input string) (JSONValue, string, bool) {
	rest := input
	rest = spaceParser(rest)

	if _, rest, ok := nullParser(rest); ok {
		return nil, rest, true
	}
	if b, rest, ok := boolParser(rest); ok {
		return b, rest, true
	}
	if n, rest, ok := numberParser(rest); ok {
		return n, rest, true
	}
	if s, rest, ok := stringParser(rest); ok {
		return s, rest, true
	}
	if o, rest, ok := objectParser(rest); ok {
		return o, rest, true
	}
	if a, rest, ok := arrayParser(rest); ok {
		return a, rest, true
	}
	return nil, input, false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filename)

	if err != nil {
		println("Error Reading File", err)
		os.Exit(1)
	}
	content := string(data)

	a, rest, ok := arrayParser(content)
	rest = spaceParser(rest)
	if len(rest) == 0 && ok {
		fmt.Println(a)
		os.Exit(0)
	}
	o, rest, ok := objectParser(content)
	rest = spaceParser(rest)
	if len(rest) == 0 && ok {
		fmt.Println(o)
		os.Exit(0)
	}
	fmt.Println("Error processing the data")

}
