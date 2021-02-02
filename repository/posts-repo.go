package repository

import (
	"github/kaji2002/entity"
)

// interface
type PostRepository interface {
	// メソッドを定義 引数は代入されるのでポインタ型にする
	Save(post *entity.Post) (*entity.Post, error)
	// 返り値がpostの配列
	FindAll() ([]entity.Post, error)
}

