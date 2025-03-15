package store

import (
	"gorm.io/gorm"
)

const (
	DefaultLimitBatchGet = 100
)

// I is ID of IEntity
type IEntity[I any] interface {
	GetId() I
}

// I type of ID, T type of model
type BaseStore[I any, T IEntity[I]] struct {
	db    *gorm.DB
	model T
}

func NewBaseStore[I any, T IEntity[I]](model T, db *gorm.DB) *BaseStore[I, T] {
	return &BaseStore[I, T]{
		db:    db,
		model: model,
	}
}
