package service

/**
 * @Author nico
 * @Date 2024-12-18
 * @File: baseService.go
 * @Description:
 */

import (
	"gin-init/model"
	"gin-init/model/types"
)

// BaseServiceInterface 基础服务接口
type BaseServiceInterface[T any] interface {
	// 基础CRUD
	Create(item *T) error
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
	Update(item *T) error
	Delete(id uint) error

	// 高级查询
	GetByCondition(condition *T) ([]T, error)
	GetByMap(condition map[string]interface{}) ([]T, error)
	GetPage(page *types.PaginationType, condition map[string]interface{}) (*types.PageResult[T], error)

	// 批量操作
	BatchCreate(items []T) error
	BatchUpdate(items []T) error
	BatchDelete(ids []uint) error

	// 软删除
	SoftDelete(id uint) error
	Restore(id uint) error
}

// BaseService 基础服务实现
type BaseService[T any] struct {
	baseModel model.BaseModelInterface[T]
}

func NewBaseService[T any]() *BaseService[T] {
	return &BaseService[T]{
		baseModel: model.NewBaseModel[T](),
	}
}

// BeforeCreate 创建前钩子
func (s *BaseService[T]) BeforeCreate(item *T) error {
	return nil
}

// AfterCreate 创建后钩子
func (s *BaseService[T]) AfterCreate(item *T) error {
	return nil
}

// BeforeUpdate 更新前钩子
func (s *BaseService[T]) BeforeUpdate(item *T) error {
	return nil
}

// AfterUpdate 更新后钩子
func (s *BaseService[T]) AfterUpdate(item *T) error {
	return nil
}

// BeforeDelete 删除前钩子
func (s *BaseService[T]) BeforeDelete(id uint) error {
	return nil
}

// AfterDelete 删除后钩子
func (s *BaseService[T]) AfterDelete(id uint) error {
	return nil
}

func (s *BaseService[T]) Create(item *T) error {
	if err := s.BeforeCreate(item); err != nil {
		return err
	}
	if err := s.baseModel.Create(item); err != nil {
		return err
	}
	return s.AfterCreate(item)
}

func (s *BaseService[T]) GetByID(id uint) (*T, error) {
	return s.baseModel.FindByID(id)
}

func (s *BaseService[T]) GetAll() ([]T, error) {
	return s.baseModel.FindAll()
}

func (s *BaseService[T]) Update(item *T) error {
	if err := s.BeforeUpdate(item); err != nil {
		return err
	}
	if err := s.baseModel.Update(item); err != nil {
		return err
	}
	return s.AfterUpdate(item)
}

func (s *BaseService[T]) Delete(id uint) error {
	if err := s.BeforeDelete(id); err != nil {
		return err
	}
	if err := s.baseModel.Delete(id); err != nil {
		return err
	}
	return s.AfterDelete(id)
}

func (s *BaseService[T]) GetByCondition(condition *T) ([]T, error) {
	return s.baseModel.FindByCondition(condition)
}

func (s *BaseService[T]) GetByMap(condition map[string]interface{}) ([]T, error) {
	return s.baseModel.FindByMap(condition)
}

func (s *BaseService[T]) GetPage(page *types.PaginationType, condition map[string]interface{}) (*types.PageResult[T], error) {
	return s.baseModel.FindPage(page, condition)
}

func (s *BaseService[T]) BatchCreate(items []T) error {
	return s.baseModel.BatchCreate(items)
}

func (s *BaseService[T]) BatchUpdate(items []T) error {
	return s.baseModel.BatchUpdate(items)
}

func (s *BaseService[T]) BatchDelete(ids []uint) error {
	return s.baseModel.BatchDelete(ids)
}

func (s *BaseService[T]) SoftDelete(id uint) error {
	return s.baseModel.SoftDelete(id)
}

func (s *BaseService[T]) Restore(id uint) error {
	return s.baseModel.Restore(id)
}
