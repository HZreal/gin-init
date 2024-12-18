package controller

/**
 * @Author nico
 * @Date 2024-12-18
 * @File: baseController.go
 * @Description:
 */

import (
	"gin-init/common/response"
	"gin-init/model/dto"
	"gin-init/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type BaseControllerInterface[T any] interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type BaseController[T any] struct {
	service service.BaseServiceInterface[T]
}

func NewBaseController[T any]() *BaseController[T] {
	return &BaseController[T]{
		service: service.NewBaseService[T](),
	}
}

// Create - 通用的创建操作
func (c *BaseController[T]) Create(ctx *gin.Context) {
	var item T
	if err := ctx.ShouldBindJSON(&item); err != nil {
		log.Println("Bind error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := c.service.Create(&item); err != nil {
		log.Println("Create error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Created successfully", "data": item})
}

// GetByID - 通用的获取单个记录
func (c *BaseController[T]) GetByID(ctx *gin.Context) {
	var body dto.QueryId

	if err := ctx.ShouldBindQuery(&body); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	item, err := c.service.GetByID(uint(body.Id))
	if err != nil {
		log.Println("Get error:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

// GetAll - 通用的获取所有记录
func (c *BaseController[T]) GetAll(ctx *gin.Context) {
	items, err := c.service.GetAll()
	if err != nil {
		log.Println("Get all error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve items"})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

// Update - 通用的更新操作
func (c *BaseController[T]) Update(ctx *gin.Context) {
	var item T
	if err := ctx.ShouldBindJSON(&item); err != nil {
		log.Println("Bind error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := c.service.Update(&item); err != nil {
		log.Println("Update error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updated successfully", "data": item})
}

// Delete - 通用的删除操作
func (c *BaseController[T]) Delete(ctx *gin.Context) {
	var body dto.BodyJsonId

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}
	if err := c.service.Delete(uint(body.Id)); err != nil {
		log.Println("Delete error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
