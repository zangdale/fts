// Package fts provides ...
package fts

import (
	"errors"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		indexDocuments = make(index)
	})
}

// Documents 文档
type Document interface {
	GetText() string
	GetID() uint64
	setID(uint64)
}

var _ Document = (*SimpleDocument)(nil)

// documents 文档夹
var documents []Document

// // NewDocuments new documents
// func NewDocuments(docs []Document) {
// 	documents = docs
// }

var (
	ErrDocumentsNull         = errors.New("documents is null")                       // 文档夹为空
	ErrDocumentsLengthLittle = errors.New("length of ids is greater than documents") // ids 长度超过 文档夹的总数
	ErrDocumentsIDOutOfScope = errors.New("ids requires data out of documents")      // ids 需要的数据不在 文档夹中
)

// GetDocuments 获取指定ID 的文档
func GetDocuments(ids []uint64) ([]Document, error) {
	if documents == nil {
		return nil, ErrDocumentsNull
	}
	dLen := len(documents)
	if dLen < len(ids) {
		return nil, ErrDocumentsLengthLittle
	}

	var res []Document
	for _, id := range ids {
		if id > uint64(dLen) {
			return nil, ErrDocumentsIDOutOfScope
		}
		res = append(res, documents[id])
	}
	return res, nil
}

var indexDocuments index

// DocumentAdd documents add docs
func DocumentAdd(docs []Document) {
	len := len(documents)
	for i := range docs {
		docs[i].setID(uint64(len + i))
		documents = append(documents, docs[i])
	}

	indexDocuments.add(docs)
}

// DocumentSearchId search ids from documents
func DocumentSearchId(text string) []uint64 {
	return indexDocuments.search(text)
}

// DocumentSearchId search document from documents
func DocumentSearch(text string) ([]Document, error) {
	return GetDocuments(DocumentSearchId(text))
}
