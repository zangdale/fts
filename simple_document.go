// Package fts provides ...
package fts

type SimpleDocument struct {
	Text string
	id   uint64
}

// NewSimpleDocument creates a new SimpleDocument
func NewSimpleDocument() *SimpleDocument {
	return &SimpleDocument{}
}

// NewSimpleDocumentWithText creates a new SimpleDocument by text
func NewSimpleDocumentWithText(text string) *SimpleDocument {
	return &SimpleDocument{Text: text}
}

func (d *SimpleDocument) SetText(text string) {
	d.Text = text
	return
}

func (d *SimpleDocument) GetText() string {
	return d.Text
}

func (d *SimpleDocument) GetID() uint64 {
	return d.id
}

func (d *SimpleDocument) SetID(id uint64) {
	d.id = id
	return
}
