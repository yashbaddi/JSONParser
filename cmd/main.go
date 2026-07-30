package main

import (
	"fmt"
	"strings"
)

func boolParser(input string) (bool, string, bool) {
	if res := strings.HasPrefix(input, "true"); res == true {
		return true, input[4:], true
	}
	if res := strings.HasPrefix(input, "false"); res == true {
		return false, input[5:], true
	}
	return false, input, false
}

func main() {
	fmt.Println(boolParser("trueafds"))
	fmt.Println(boolParser("falseafds"))
	fmt.Println(boolParser("faleafds"))

}
