// Package fts provides ...
package fts

// index is an inverted index. It maps tokens to document IDs.
type index map[string][]uint64

// add adds documents to the index.
func (idx index) add(docs []Document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.GetText()) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.GetID() {
				// Don't add same ID twice.
				continue
			}
			idx[token] = append(ids, doc.GetID())
		}
	}
}

// intersection returns the set intersection between a and b.
// a and b have to be sorted in ascending order and contain no duplicates.
func intersection(a []uint64, b []uint64) []uint64 {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]uint64, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// search queries the index for the given text.
func (idx index) search(text string) []uint64 {
	var r []uint64
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return r
}
