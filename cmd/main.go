package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func nullParser(input string) (any, string, bool) {
	rest := input
	if res := strings.HasPrefix(rest, "null"); res == true {
		return nil, rest[4:], true
	}
	return nil, input, false
}

func boolParser(input string) (bool, string, bool) {
	rest := input
	if res := strings.HasPrefix(rest, "true"); res == true {
		return true, rest[4:], true
	}
	if res := strings.HasPrefix(rest, "false"); res == true {
		return false, rest[5:], true
	}
	return false, input, false
}

func numberParser(input string) (float64, string, bool) {
	rest := input
	re := regexp.MustCompile(`^-?([1-9](\d)*|0)(\.(\d)+)?([eE][+-]?(\d)+)?`)

	match := re.FindString(rest)
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

	for i > len(rest) && rest[i] != '"' {
		if err := builder.WriteByte(rest[i]); err != nil {
			return "", input, false
		}
		i++
	}

	if i == len(rest) {
		return "", input, false
	}

	return builder.String(), rest[i+1:], true
}

func objectParser(input string) (map[string]any, string, bool) {
	rest := input

	if !strings.HasPrefix(rest, "{") {
		return nil, input, false

	}

	data := make(map[string]any)
	rest = spaceParser(rest[1:])

	for !strings.HasPrefix(rest, "}") {
		rest = spaceParser(rest)

		key, rest, ok := stringParser(rest)
		if !ok {
			return nil, input, false
		}

		rest = spaceParser(rest)

		if !strings.HasPrefix(rest, ":") {
			return nil, input, false
		}

		value, rest, ok := valueParser(rest[1:])
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

type JSONValue = any

func valueParser(input string) (JSONValue, string, bool) {
	rest := input
	rest = spaceParser(rest)

	if nil, rest, ok := nullParser(rest); ok {
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
	return nil, input, false
}

func main() {
	// fmt.Println(boolParser("trueafds"))
	// fmt.Println(boolParser("falseafds"))
	// fmt.Println(boolParser("faleafds"))
	fmt.Println(numberParser("24fdsaf"))
	fmt.Println(stringParser("\""))
	a := make(map[string]any)
	a["asdsa"] = 0
	v, _ := json.Marshal(a)
	fmt.Println(string(v))
	fmt.Println(objectParser(string(v)))

}
