package controller

/**
 * @Author nico
 * @Date 2024-12-18
 * @File: baseController.go
 * @Description:
 */

import (
	"gin-init/common/response"
	"gin-init/model/types"
	"gin-init/service"

	"github.com/gin-gonic/gin"
)

// BaseControllerInterface 基础控制器接口
type BaseControllerInterface[T any] interface {
	// 基础CRUD
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)

	// 高级查询
	GetByCondition(ctx *gin.Context)
	GetByMap(ctx *gin.Context)
	GetPage(ctx *gin.Context)

	// 批量操作
	BatchCreate(ctx *gin.Context)
	BatchUpdate(ctx *gin.Context)
	BatchDelete(ctx *gin.Context)

	// 软删除
	SoftDelete(ctx *gin.Context)
	Restore(ctx *gin.Context)
}

// BaseController 基础控制器实现
type BaseController[T any] struct {
	service service.BaseServiceInterface[T]
}

func NewBaseController[T any]() *BaseController[T] {
	return &BaseController[T]{
		service: service.NewBaseService[T](),
	}
}

// Create 创建记录
func (c *BaseController[T]) Create(ctx *gin.Context) {
	var item T
	if err := ctx.ShouldBindJSON(&item); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	if err := c.service.Create(&item); err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, item)
}

// GetByID 获取单个记录
func (c *BaseController[T]) GetByID(ctx *gin.Context) {
	var body types.QueryId
	if err := ctx.ShouldBindQuery(&body); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	item, err := c.service.GetByID(body.Id)
	if err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, item)
}

// GetAll 获取所有记录
func (c *BaseController[T]) GetAll(ctx *gin.Context) {
	items, err := c.service.GetAll()
	if err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, items)
}

// Update 更新记录
func (c *BaseController[T]) Update(ctx *gin.Context) {
	var item T
	if err := ctx.ShouldBindJSON(&item); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	if err := c.service.Update(&item); err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, item)
}

// Delete 删除记录
func (c *BaseController[T]) Delete(ctx *gin.Context) {
	var body types.BodyJsonId
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	if err := c.service.Delete(uint(body.Id)); err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithoutData(ctx)
}

// GetByCondition 条件查询
func (c *BaseController[T]) GetByCondition(ctx *gin.Context) {
	var condition T
	if err := ctx.ShouldBindJSON(&condition); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	items, err := c.service.GetByCondition(&condition)
	if err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, items)
}

// GetByMap Map条件查询
func (c *BaseController[T]) GetByMap(ctx *gin.Context) {
	var condition map[string]interface{}
	if err := ctx.ShouldBindJSON(&condition); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	items, err := c.service.GetByMap(condition)
	if err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, items)
}

// GetPage 分页查询
func (c *BaseController[T]) GetPage(ctx *gin.Context) {
	var query struct {
		types.PaginationType
		Condition map[string]interface{} `json:"condition"`
	}

	if err := ctx.ShouldBindJSON(&query); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	result, err := c.service.GetPage(&query.PaginationType, query.Condition)
	if err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, result)
}

// BatchCreate 批量创建
func (c *BaseController[T]) BatchCreate(ctx *gin.Context) {
	var items []T
	if err := ctx.ShouldBindJSON(&items); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	if err := c.service.BatchCreate(items); err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, items)
}

// BatchUpdate 批量更新
func (c *BaseController[T]) BatchUpdate(ctx *gin.Context) {
	var items []T
	if err := ctx.ShouldBindJSON(&items); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	if err := c.service.BatchUpdate(items); err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, items)
}

// BatchDelete 批量删除
func (c *BaseController[T]) BatchDelete(ctx *gin.Context) {
	var body struct {
		Ids []uint `json:"ids"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	if err := c.service.BatchDelete(body.Ids); err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithoutData(ctx)
}

// SoftDelete 软删除
func (c *BaseController[T]) SoftDelete(ctx *gin.Context) {
	var body types.BodyJsonId
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	if err := c.service.SoftDelete(uint(body.Id)); err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithoutData(ctx)
}

// Restore 恢复删除
func (c *BaseController[T]) Restore(ctx *gin.Context) {
	var body types.BodyJsonId
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	if err := c.service.Restore(uint(body.Id)); err != nil {
		response.FailedWithMsg(ctx, err.Error())
		return
	}

	response.SuccessWithoutData(ctx)
}
