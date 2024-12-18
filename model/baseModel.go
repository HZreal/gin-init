package model

/**
 * @Author nico
 * @Date 2024-12-18
 * @File: baseModel.go
 * @Description:
 */

import (
	"gin-init/core/initialize/database"
	"gorm.io/gorm"
)

type BaseRepoInterface[T any] interface {
	Create(item *T) error
	FindByID(id uint) (*T, error)
	FindAll() ([]T, error)
	Update(item *T) error
	Delete(id uint) error
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{
		db: database.DB,
	}
}

func (r *BaseRepository[T]) Create(item *T) error {
	return r.db.Create(item).Error
}

func (r *BaseRepository[T]) FindByID(id uint) (*T, error) {
	var item T
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var items []T
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BaseRepository[T]) Update(item *T) error {
	return r.db.Save(item).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	var item T
	return r.db.Delete(item, id).Error
}
