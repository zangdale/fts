// Package fts provides ...
package fts

import "strings"

// FilterFunc 过滤函数
type FilterFunc func(tokens []string) []string

var filterFuncs []FilterFunc

// loadFilter tokens filter
func loadFilter(tokens []string) []string {
	if len(filterFuncs) == 0 {
		return tokens
	}

	var res = tokens
	for _, f := range filterFuncs {
		if res == nil {
			return res
		}
		res = f(res)
	}
	return res
}

// AddFilter add filter function
func AddFilter(f ...FilterFunc) {
	filterFuncs = append(filterFuncs, f...)
}

// TrimSpaceFilter returns a slice of no Space.
var TrimSpaceFilter = func(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if len(token) == 0 {
			continue
		}
		r = append(r, token)
	}
	return r
}

// lowercaseFilter returns a slice of tokens normalized to lower case.
var LowercaseFilter = func(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

// stopwordFilter returns a slice of tokens with stop words removed.
var StopwordFilter = func(tokens []string) []string {
	var stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}
