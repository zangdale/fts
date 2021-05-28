package main

type Docs struct {
	Text   string
	id     uint64
	Action func()
}

// NewDocs creates a new Docs
func NewDocs() *Docs {
	return &Docs{}
}

func (d *Docs) GetText() string {
	return d.Text
}
func (d *Docs) GetID() uint64 {
	return d.id
}
func (d *Docs) SetID(id uint64) {
	d.id = id
}
