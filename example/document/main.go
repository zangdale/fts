package main

import (
	"fmt"

	"github.com/getbuguai/fts"
)

var _ fts.Document = (*Docs)(nil)

func main() {
	fmt.Println("Hello BuGuai !!! ")
	var strs = []string{"你好", "你们好", "你什么", "好好什么", "说是拉开"}
	fts.DefaultSplitString = fts.ChaineseSplitString

	fts.DocumentAdd(sliceStrToSimpleDocument(strs))

	ds, err := fts.DocumentSearch("你好")
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
			Text: v,
			Action: func(t string) func() {
				return func() {
					println(t)
				}
			}(v),
		}
	}
	return res
}
