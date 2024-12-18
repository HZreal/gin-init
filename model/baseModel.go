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

type BaseModelInterface[T any] interface {
	Create(item *T) error
	FindByID(id uint) (*T, error)
	FindAll() ([]T, error)
	Update(item *T) error
	Delete(id uint) error
}

type BaseModel[T any] struct {
	db *gorm.DB
}

func NewBaseModel[T any]() *BaseModel[T] {
	return &BaseModel[T]{
		db: database.DB,
	}
}

func (r *BaseModel[T]) Create(item *T) error {
	result := r.db.Create(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *BaseModel[T]) FindByID(id uint) (*T, error) {
	var item T
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *BaseModel[T]) FindAll() ([]T, error) {
	var items []T
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BaseModel[T]) Update(item *T) error {
	return r.db.Save(item).Error
}

func (r *BaseModel[T]) Delete(id uint) error {
	var item T
	return r.db.Delete(item, id).Error
}
