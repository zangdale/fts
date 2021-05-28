// Package fts provides ...
package fts

type SimpleDocument struct {
	Text   string
	id     uint64
	Action func()
}

func (d *SimpleDocument) GetText() string {
	return d.Text
}

func (d *SimpleDocument) GetID() uint64 {
	return d.id
}

func (d *SimpleDocument) setID(id uint64) {
	d.id = id
	return
}
