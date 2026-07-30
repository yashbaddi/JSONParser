package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func nullParser(input string) (string, bool) {
	if res := strings.HasPrefix(input, "null"); res == true {
		return input[4:], true
	}
	return input, false
}

func boolParser(input string) (bool, string, bool) {
	if res := strings.HasPrefix(input, "true"); res == true {
		return true, input[4:], true
	}
	if res := strings.HasPrefix(input, "false"); res == true {
		return false, input[5:], true
	}
	return false, input, false
}

func numberParser(input string) (float64, string, bool) {
	re := regexp.MustCompile(`^-?([1-9](\d)*|0)(\.(\d)+)?([eE][+-]?(\d)+)?`)

	match := re.FindString(input)
	if match == "" {
		return 0, input, false
	}

	num, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, input, false
	}
	return num, input[len(match):], true
}

func stringParser(input string) (string, string, bool) {
	var builder strings.Builder
	if input[0] == '"' {
		i := 0
		for input[i] != '"' {
			if err := builder.WriteByte(input[i]); err != nil {
				return "", input, false
			}
		}
		return builder.String(), input[i:], true
	}
	return "", input, false
}

func valueParser(input string) (any, string, bool) {
	input = strings.TrimSpace(input)

	if rest, ok := nullParser(input); ok {
		return nil, rest, true
	}
	if b, rest, ok := boolParser(input); ok {
		return b, rest, true
	}
	if n, rest, ok := numberParser(input); ok {
		return n, rest, true
	}
	if s, rest, ok := boolParser(input); ok {
		return s, rest, true
	}
	return nil, input, false
}

func main() {
	// fmt.Println(boolParser("trueafds"))
	// fmt.Println(boolParser("falseafds"))
	// fmt.Println(boolParser("faleafds"))
	fmt.Println(numberParser("24fdsaf"))
	fmt.Println(stringParser("\""))

}
