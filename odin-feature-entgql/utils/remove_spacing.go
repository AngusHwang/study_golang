package utils

import "strings"

func RemoveSpacing(word string) string {
	result := strings.ReplaceAll(word, " ", "")

	return result
}
