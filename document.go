// Package fts provides ...
package fts

import "errors"

// Documents 文档
type Document interface {
	GetText() string
	GetID() uint64
}

// Documents 文档夹
var Documents []Document

var (
	ErrDocumentsNull         = errors.New("documents is null")                       // 文档夹为空
	ErrDocumentsLengthLittle = errors.New("length of ids is greater than documents") // ids 长度超过 文档夹的总数
	ErrDocumentsIDOutOfScope = errors.New("ids requires data out of documents")      // ids 需要的数据不在 文档夹中
)

// GetDocuments 获取指定ID 的文档
func GetDocuments(ids []uint64) ([]Document, error) {
	if Documents == nil {
		return nil, ErrDocumentsNull
	}
	dLen := len(Documents)
	if dLen < len(ids) {
		return nil, ErrDocumentsLengthLittle
	}

	var res []Document
	for _, id := range ids {
		if id > uint64(dLen) {
			return nil, ErrDocumentsIDOutOfScope
		}
		res = append(res, Documents[id])
	}
	return res, nil
}
