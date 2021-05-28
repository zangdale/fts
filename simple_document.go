// Package fts provides ...
package fts

type SimpleDocument struct {
	Text   string
	ID     uint64
	Action func()
}

func (d *SimpleDocument) GetText() string {
	return d.Text
}

func (d *SimpleDocument) GetID() uint64 {
	return d.ID
}
