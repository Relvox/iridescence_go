package utils

import (
	"fmt"
	"strings"
)

func Strings[T fmt.Stringer](slice []T) []string {
	var result []string = make([]string, len(slice))
	for i := 0; i < len(result); i++ {
		result[i] = fmt.Sprint(slice[i])
	}
	return result
}

func ToCamelCase(input string) string {
	result := strings.ReplaceAll(input, " ", "")
	return result
}
