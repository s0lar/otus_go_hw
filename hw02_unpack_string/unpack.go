package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(source string) (string, error) {
	arrRunes := []rune(source)
	var result strings.Builder

	for i, currRune := range arrRunes {
		var nextRune rune
		currChar := string(currRune)
		nextChar := string(nextRune)
		currRuneIsDigit := unicode.IsDigit(currRune)
		nextRuneIsDigit := false

		if i < len(arrRunes)-1 {
			nextRune = arrRunes[i+1]
			nextChar = string(nextRune)
			nextRuneIsDigit = unicode.IsDigit(nextRune)
		}

		if currRuneIsDigit && i == 0 {
			return "", ErrInvalidString
		}

		if currRuneIsDigit && nextRuneIsDigit {
			return "", ErrInvalidString
		}

		if currRuneIsDigit {
			continue
		}

		repeat := 1
		if nextRuneIsDigit {
			repeat, _ = strconv.Atoi(nextChar)
		}

		result.WriteString(strings.Repeat(currChar, repeat))
	}

	if source != "" && result.String() == "" {
		return "", ErrInvalidString
	}

	return result.String(), nil
}
