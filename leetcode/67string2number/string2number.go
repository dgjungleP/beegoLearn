package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stat int
type CharType int

const INT_MAX = 2147483647
const INT_MIN = -2147483648
const (
	EMPTY Stat = iota
	SIGN
	NUMBER
	STRING
)
const (
	IS_EMPTY CharType = iota
	IS_SIGN
	IS_NUMBER
	IS_STRING
)

func strToInt(str string) int {
	ruleMap := map[Stat]map[CharType]Stat{
		EMPTY: map[CharType]Stat{
			IS_EMPTY:  EMPTY,
			IS_SIGN:   SIGN,
			IS_NUMBER: NUMBER,
		},
		SIGN: map[CharType]Stat{
			IS_NUMBER: NUMBER,
		},
		NUMBER: map[CharType]Stat{
			IS_NUMBER: NUMBER,
		},
	}
	state := EMPTY
	for index, character := range str {
		charType := getType(character)
		if value, ok := ruleMap[state][charType]; !ok {
			result, _ := strconv.Atoi(strings.Trim(str[:index], " "))
			if result < INT_MIN {
				return INT_MIN
			} else if result > INT_MAX {
				return INT_MAX
			}
			return result
		} else {
			state = value
		}
	}
	result, _ := strconv.Atoi(strings.Trim(str, " "))
	if result < INT_MIN {
		return INT_MIN
	} else if result > INT_MAX {
		return INT_MAX
	}
	return result
}

func getType(character rune) CharType {
	switch character {
	case ' ':
		return IS_EMPTY
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return IS_NUMBER
	case '-', '+':
		return IS_SIGN
	default:
		return IS_STRING
	}
}

func main() {
	value := strToInt("-91283472332")
	fmt.Print(value)
}
