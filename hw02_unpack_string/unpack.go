package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const (
	Backslash rune = '\\'

	TypeDigit = iota
	TypeSlash
	TypeSymbl
	TypeStart
)

func ActionForDigit(predType int, currSymbl, predSymbl rune, lastS bool, answer []rune) (int, []rune, error) {
	switch predType {
	case TypeDigit:
		return 0, nil, ErrInvalidString
	case TypeStart:
		return 0, nil, ErrInvalidString

	case TypeSlash:
		if lastS {
			answer = append(answer, currSymbl)
		}
		predType = TypeSymbl

	default:
		currS, predS := string(currSymbl), string(predSymbl)
		count, err := strconv.Atoi(currS)
		if err != nil {
			return 0, nil, err
		}
		if count > 0 {
			answer = append(answer, []rune(strings.Repeat(predS, count))...)
		}
		predType = TypeDigit
	}

	return predType, answer, nil
}

func ActionForSlash(predType int, currSymbl, predSymbl rune, lastS bool, answer []rune) (int, []rune, error) {
	switch predType {
	case TypeSlash:
		if lastS {
			answer = append(answer, currSymbl)
		}
		predType = TypeSymbl

	case TypeSymbl:
		answer = append(answer, predSymbl)
		predType = TypeSlash
	default:
		predType = TypeSlash
	}
	return predType, answer, nil
}

func ActionForSymbl(predType int, currSymbl, predSymbl rune, lastS bool, answer []rune) (int, []rune, error) {
	switch predType {
	case TypeSlash:
		return 0, []rune{}, ErrInvalidString

	case TypeSymbl:
		answer = append(answer, predSymbl)
	}
	predType = TypeSymbl
	if lastS {
		answer = append(answer, currSymbl)
	}
	return predType, answer, nil
}

func Unpack(s string) (string, error) {
	str := []rune(s)
	answer := make([]rune, 0, len(str))
	var err error
	var predSymbl rune

	predType := TypeStart

	for i, symbl := range str {
		switch {
		case unicode.IsDigit(symbl):

			predType, answer, err = ActionForDigit(predType, symbl, predSymbl, i == len(str)-1, answer)
			if err != nil {
				return "", err
			}

		case symbl == Backslash:
			predType, answer, err = ActionForSlash(predType, symbl, predSymbl, i == len(str)-1, answer)
			if err != nil {
				return "", err
			}

		default:
			predType, answer, err = ActionForSymbl(predType, symbl, predSymbl, i == len(str)-1, answer)
			if err != nil {
				return "", err
			}
		}
		predSymbl = symbl
	}

	return string(answer), nil
}
