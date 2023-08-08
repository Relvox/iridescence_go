package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func Strings[T fmt.Stringer](slice []T) []string {
	var result []string = make([]string, len(slice))
	for i := 0; i < len(result); i++ {
		result[i] = fmt.Sprint(slice[i])
	}
	return result
}

func Indented[T ~string](indent T, strings ...T) []T {
	result := make([]T, len(strings))
	for i := 0; i < len(strings); i++ {
		result[i] = T(fmt.Sprintf("%s%s", indent, strings[i]))
	}
	return result
}

type StringCase int

const (
	// camelCaseText
	CamelCase StringCase = iota
	// PascalCaseText
	PascalCase
	// snake_case_text
	SnakeCase
	// Title Case Text
	TitleCase
)

func (c StringCase) Render(words []string) string {
	sb := strings.Builder{}
	switch c {
	case CamelCase:
		for i, w := range words {
			if i == 0 {
				sb.WriteString(strings.ToLower(w))
				continue
			}
			sb.WriteString(strings.ToUpper(w[:1]))
			sb.WriteString(strings.ToLower(w[1:]))
		}
	case PascalCase:
		for _, w := range words {
			sb.WriteString(strings.ToUpper(w[:1]))
			sb.WriteString(strings.ToLower(w[1:]))
		}
	case SnakeCase:
		for i, w := range words {
			sb.WriteString(strings.ToLower(w))
			if i == len(words)-1 {
				continue
			}
			sb.WriteString("_")
		}
	case TitleCase:
		for i, w := range words {
			sb.WriteString(strings.ToUpper(w[:1]))
			sb.WriteString(strings.ToLower(w[1:]))
			if i == len(words)-1 {
				continue
			}
			sb.WriteString(" ")
		}
	default:
		panic("unknown case")
	}
	return sb.String()
}

func (c StringCase) Parse(phrase string) []string {
	var result []string
	switch c {
	case PascalCase, CamelCase:
		last := 0
		for i, r := range phrase {
			if i == 0 || !unicode.IsUpper(r) {
				continue
			}
			result = append(result, strings.ToLower(phrase[last:i]))
			last = i
		}

	case SnakeCase:
		result = strings.Split(phrase, "_")

	case TitleCase:
		result = strings.Split(phrase, " ")
		for i, v := range result {
			result[i] = strings.ToLower(v)
		}

	default:
		panic("unknown case")
	}
	return result
}

func Transcase(input string, from, to StringCase) string {
	words := from.Parse(input)
	return to.Render(words)
}
