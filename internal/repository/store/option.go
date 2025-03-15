package store

import (
	"gorm.io/gorm"
)

type Option func(*gorm.DB) *gorm.DB

func Condition(query any, args ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func Limit(limit int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func OrderBy(value any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(value)
	}
}

func Offset(offset int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}

func Preload(column string, args ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(column, args...)
	}
}

func Select(query any, args ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(query, args...)
	}
}
