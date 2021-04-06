package DDATM_ES_0001

import (
	"strings"
)

func Exec(disorder string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "ñ", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var disorderArrray []string
	var orderArray []string

	for i := 0; i < len(disorder); i++ {
		disorderArrray = append(disorderArrray, disorder[i:i+1])
	}

	for _, compareDisorder := range disorderArrray {
		orderArray = append(orderArray, returnNumberAlphabet(compareDisorder, alphabet))
	}

	orderString := strings.Join(orderArray, "")

	return orderString
}

func returnNumberAlphabet(compare string, alphabet []string) string {
	for j, alalphabetCompare := range alphabet {
		if strings.ToLower(compare) == alalphabetCompare {
			if strings.ToLower(compare) == alphabet[26] {
				return validateLetter(compare, alphabet[0])
			}
			return validateLetter(compare, alphabet[j+1])
		}
	}

	return compare
}

func validateLetter(compare ,alphabetAssert string) string {
	compareUpper := strings.ToUpper(compare)

	if compareUpper == compare {
		return strings.ToUpper(alphabetAssert)
	}

	return alphabetAssert
}
