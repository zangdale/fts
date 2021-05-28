// Package fts provides ...
package fts

import (
	"strings"
)

// DefaultSplitString 默认拆分规则函数
var DefaultSplitString = EnglishSplitString

// EnglishSplitString splits a string
func EnglishSplitString(s string) []string {

	// return strings.FieldsFunc(s, func(r rune) bool {
	// 	// Split on any character that is not a letter or a number.
	// 	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	// })
	return strings.Split(s, " ")
}

// ChaineseSplitString splits a string
func ChaineseSplitString(s string) []string {

	// return strings.FieldsFunc(s, func(r rune) bool {
	// 	// Split on any character that is not a letter or a number.
	// 	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	// })
	var res []string
	for _, v := range s {
		if string(v) == "" {
			continue
		}
		res = append(res, string(v))
	}
	return res
}

// analyze analyzes the text and returns a slice of tokens.
func analyze(text string) []string {
	tokens := DefaultSplitString(text)

	tokens = loadFilter(tokens)
	return tokens
}
