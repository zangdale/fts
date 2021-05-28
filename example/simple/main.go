// Package main provides ...
package main

import (
	"fmt"

	"github.com/getbuguai/fts"
)

type Docs struct {
	fts.SimpleDocument
}

func main() {
	fmt.Println("Hello BuGuai !!! ")
	var strs = []string{"你好", "射门申请", "看看看么", "看看", "说是拉开"}
	fts.DefaultSplitString = fts.ChaineseSplitString

	fts.DocumentAdd(sliceStrToSimpleDocument(strs))

	// TODO: 中文叠词错误,会匹配 [看看看么,看看]
	ds, err := fts.DocumentSearch("看看看看")
	if err != nil {
		fmt.Println("DocumentSearch", err)
		return
	}
	for _, d := range ds {
		s := d.(*Docs)
		fmt.Println(s.GetText())
	}
}

func sliceStrToSimpleDocument(slicesStrs []string) []fts.Document {
	res := make([]fts.Document, len(slicesStrs))
	for i, v := range slicesStrs {
		res[i] = &Docs{
			SimpleDocument: fts.SimpleDocument{
				Text: v,
			},
		}
	}
	return res
}
