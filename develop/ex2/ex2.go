package main

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	REGEX_NUMBERS string = "[0-9]*"
	REGEX_SYMBOLS string = "[a-zA-Z]"
)

func unpackingString(str string) (string, error) {
	if str == "" {
		return str, nil
	}
	reNumber := regexp.MustCompile(REGEX_NUMBERS)
	reSymbol := regexp.MustCompile(REGEX_SYMBOLS)
	arrNumber := reNumber.FindAllString(str, -1)
	arrString := reSymbol.FindAllString(str, -1)
	if len(arrNumber) > len(str) {
		return str, nil
	} else if len(arrString) > len(str) {
		return "", fmt.Errorf("Incorrect string")
	}
	var result string = ""
	for i, val := range arrString {
		number, _ := strconv.Atoi(arrNumber[len(val)+i])
		if number == 0 {
			result += val
			continue
		}
		for j := 0; j < number; j++ {
			result += val
		}
	}
	return result, nil
}
