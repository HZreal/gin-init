package service

/**
 * @Author nico
 * @Date 2024-12-18
 * @File: baseService.go
 * @Description:
 */

import (
	"gin-init/model"
)

type BaseServiceInterface[T any] interface {
	Create(item *T) error
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
	Update(item *T) error
	Delete(id uint) error
}

type BaseService[T any] struct {
	repository model.BaseModelInterface[T]
}

func NewBaseService[T any]() *BaseService[T] {
	return &BaseService[T]{
		repository: model.NewBaseModel[T](),
	}
}

func (s *BaseService[T]) Create(item *T) error {
	return s.repository.Create(item)
}

func (s *BaseService[T]) GetByID(id uint) (*T, error) {
	return s.repository.FindByID(id)
}

func (s *BaseService[T]) GetAll() ([]T, error) {
	return s.repository.FindAll()
}

func (s *BaseService[T]) Update(item *T) error {
	return s.repository.Update(item)
}

func (s *BaseService[T]) Delete(id uint) error {
	return s.repository.Delete(id)
}
