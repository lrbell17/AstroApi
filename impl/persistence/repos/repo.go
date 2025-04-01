package repos

import (
	"github.com/lrbell17/astroapi/impl/persistence/dao"
)

type AstroRepo[T dao.AstroDAO] interface {
	GetById(id uint) (T, error)
	Insert(dao T) (T, error)
	BatchInsert(batch []T) (int, error)
	GetAll() ([]T, error)
}
