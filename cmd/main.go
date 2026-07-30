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

func main() {
	// fmt.Println(boolParser("trueafds"))
	// fmt.Println(boolParser("falseafds"))
	// fmt.Println(boolParser("faleafds"))
	fmt.Println(numberParser("24fdsaf"))

}
