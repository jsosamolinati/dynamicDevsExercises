package DDATM_ES_0001

import (
	"strings"
)

func Exec(sentence string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "ñ", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var sentenceSeparate []string
	var sentenceNeat []string

	for i := 0; i < len(sentence); i++ {
		sentenceSeparate = append(sentenceSeparate, sentence[i:i+1])
	}

	for _, compareDisorder := range sentenceSeparate {
		sentenceNeat = append(sentenceNeat, returnNumberAlphabet(compareDisorder, alphabet))
	}

	orderSentence := strings.Join(sentenceNeat, "")

	return orderSentence
}

func returnNumberAlphabet(compare string, alphabet []string) string {
	for j, alalphabetCompare := range alphabet {
		if strings.ToLower(compare) == alalphabetCompare {
			if strings.ToLower(compare) == alphabet[26] {
				return validateCapitalLetter(compare, alphabet[0])
			}
			return validateCapitalLetter(compare, alphabet[j+1])
		}
	}

	return compare
}

func validateCapitalLetter(compare ,alphabetAssert string) string {
	compareUpper := strings.ToUpper(compare)

	if compareUpper == compare {
		return strings.ToUpper(alphabetAssert)
	}

	return alphabetAssert
}
