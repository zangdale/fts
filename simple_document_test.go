// Package fts provides ...
package fts

import (
	"testing"
)

// Test .
func TestSimpleDocument(t *testing.T) {
	// var strs = []string{"test one  one", "hello  name", "test two", "test   hello name", "test One"}
	var strs = []string{"你好", "射门申请", "看看看么", "看看", "说是拉开"}
	DefaultSplitString = ChaineseSplitString

	DocumentAdd(sliceStrToSimpleDocument(strs))
	// DocumentAdd(sliceStrToSimpleDocument(strs))
	// AddFilter(func(tokens []string) []string {
	// 	r := make([]string, len(tokens))
	// 	for i, token := range tokens {
	// 		low := strings.ToLower(token)
	// 		if low == "one" {
	// 			return nil
	// 		}
	// 		r[i] = low
	// 	}
	// 	return r
	// })
	// AddFilter(TrimSpaceFilter)

	// TODO: 中文叠词错误,会匹配 [看看看么,看看]
	ds, err := DocumentSearch("看看看看")
	if err != nil {
		t.Fatal("DocumentSearch", err)
	}
	for _, d := range ds {
		s := d.(*SimpleDocument)
		s.Action()
	}

}

func sliceStrToSimpleDocument(slicesStrs []string) []Document {
	res := make([]Document, len(slicesStrs))
	for i, v := range slicesStrs {
		res[i] = &SimpleDocument{
			Text: v,
			Action: func(text string) func() {
				return func() {
					println(text)
				}
			}(v),
		}
	}
	return res
}
