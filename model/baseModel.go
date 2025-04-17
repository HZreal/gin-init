package model

/**
 * @Author nico
 * @Date 2024-12-18
 * @File: baseModel.go
 * @Description:
 */

import (
	"gin-init/core/initialize/database"
	"gin-init/model/types"
	"time"

	"gorm.io/gorm"
)

// BaseModelInterface 基础模型接口
type BaseModelInterface[T any] interface {
	// 基础CRUD
	Create(item *T) error
	FindByID(id uint) (*T, error)
	FindAll() ([]T, error)
	Update(item *T) error
	Delete(id uint) error

	// 高级查询
	FindByCondition(condition *T) ([]T, error)
	FindByMap(condition map[string]interface{}) ([]T, error)
	FindPage(page *types.PaginationType, condition map[string]interface{}) (*types.PageResult[T], error)

	// 批量操作
	BatchCreate(items []T) error
	BatchUpdate(items []T) error
	BatchDelete(ids []uint) error

	// 软删除
	SoftDelete(id uint) error
	Restore(id uint) error

	// 自定义查询
	GetDB() *gorm.DB
}

// BaseEntity 基础实体结构
type BaseEntity struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

// BaseModel 基础模型实现
type BaseModel[T any] struct {
	db *gorm.DB
}

func NewBaseModel[T any]() *BaseModel[T] {
	return &BaseModel[T]{
		db: database.DB,
	}
}

func (r *BaseModel[T]) Create(item *T) error {
	return r.db.Create(item).Error
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
	return r.db.Delete(&item, id).Error
}

func (r *BaseModel[T]) FindByCondition(condition *T) ([]T, error) {
	var items []T
	if err := r.db.Where(condition).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BaseModel[T]) FindByMap(condition map[string]interface{}) ([]T, error) {
	var items []T
	if err := r.db.Where(condition).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BaseModel[T]) FindPage(page *types.PaginationType, condition map[string]interface{}) (*types.PageResult[T], error) {
	var items []T
	var total int64

	// 计算总数
	query := r.db.Model(new(T))
	if len(condition) > 0 {
		query = query.Where(condition)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// 查询数据
	offset := (page.Page - 1) * page.PageSize
	if err := query.Offset(offset).Limit(page.PageSize).Find(&items).Error; err != nil {
		return nil, err
	}

	return &types.PageResult[T]{
		Items:    items,
		Total:    total,
		Page:     page.Page,
		PageSize: page.PageSize,
	}, nil
}

func (r *BaseModel[T]) BatchCreate(items []T) error {
	return r.db.Create(&items).Error
}

func (r *BaseModel[T]) BatchUpdate(items []T) error {
	for _, item := range items {
		if err := r.db.Save(&item).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *BaseModel[T]) BatchDelete(ids []uint) error {
	var item T
	return r.db.Delete(&item, ids).Error
}

func (r *BaseModel[T]) SoftDelete(id uint) error {
	var item T
	return r.db.Delete(&item, id).Error
}

func (r *BaseModel[T]) Restore(id uint) error {
	var item T
	return r.db.Unscoped().Model(&item).Where("id = ?", id).Update("deleted_at", nil).Error
}

func (r *BaseModel[T]) GetDB() *gorm.DB {
	return r.db
}
