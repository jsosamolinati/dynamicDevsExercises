package exercises

import (
	"strings"
)

func Execute(disorder string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "ñ", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	disorderArrray := strings.Split(disorder, ",")
	var orderArray []string

	for i, compareDisorder := range disorderArrray{
		for j, alalphabetCompare := range alphabet{
			if (compareDisorder == alalphabetCompare) {
				orderArray[i] = string(rune(j + 1))
			}
			{
				orderArray[i] = compareDisorder
			}
		}
	}

	orderString := strings.Join(orderArray, "")

	return orderString
}