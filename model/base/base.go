package ddd

/**
 * @Author elastic·H
 * @Date 2024-10-23
 * @File: base.go
 * @Description:
 */

import (
	"context"
)

type EntityType interface {
}

type BaseModelInterface[T any] interface {
	Insert(ctx context.Context, data *T) (*T, error)
	FindOne(ctx context.Context, id int64) (*T, error)
	Update(ctx context.Context, data *T) error
	Delete(ctx context.Context, id int64) error
}

type BaseRepo[T any] struct {
	conn  string
	table string
}

func NewBaseRepo[T any]() *BaseRepo[T] {
	return &BaseRepo[T]{conn: "", table: ""}
}

func (m *BaseRepo[T]) Insert(ctx context.Context, data *T) (*T, error) {
	return data, nil
}

func (m *BaseRepo[T]) FindOne(ctx context.Context, id int64) (*T, error) {
	var result T
	return &result, nil
}

func (m *BaseRepo[T]) Update(ctx context.Context, data *T) error {
	return nil
}

func (m *BaseRepo[T]) Delete(ctx context.Context, id int64) error {
	return nil
}

// ///////////////////////////////////////////////////////////////////////////////

type UserRepoInterface[T any] interface {
	BaseModelInterface[T]

	f1(ctx context.Context, id int) *T
}

type UserEntity struct {
	Id   int
	Name string
}

type UserRepo[T any] struct {
	*BaseRepo[T]
}

func NewUserRepo() UserRepoInterface[UserEntity] {
	return &UserRepo[UserEntity]{
		BaseRepo: NewBaseRepo[UserEntity](),
	}
}

func (m *UserRepo[T]) f1(ctx context.Context, id int) *T {
	var result T
	return &result
}

// ///////////////////////////////////////////////////////////////////////////////

type OrderRepoInterface[T any] interface {
	BaseModelInterface[T]

	f2(ctx context.Context, id int) bool
}

type OrderEntity struct {
	Id   int
	DESC string
}

type OrderRepo[T any] struct {
	*BaseRepo[T]
}

func NewOrderRepo() OrderRepoInterface[OrderEntity] {
	return &OrderRepo[OrderEntity]{
		BaseRepo: NewBaseRepo[OrderEntity](),
	}
}

func (m *OrderRepo[T]) f2(ctx context.Context, id int) bool {
	return true
}

// ///////////////////////////////////////////////////////////////////////////////

type UserService[T any] struct {
	U UserRepoInterface[T]
}

func NewUserService() *UserService[UserEntity] {
	return &UserService[UserEntity]{
		U: NewUserRepo(),
	}
}

// ///////////////////////////////////////////////////////////////////////////////

type OrderService[T any] struct {
	O OrderRepoInterface[T]
}

func NewOrderService() *OrderService[OrderEntity] {
	return &OrderService[OrderEntity]{
		O: NewOrderRepo(),
	}
}

// ///////////////////////////////////////////////////////////////////////////////

func main() {
	userService := NewUserService()
	userService.U.FindOne(context.Background(), 1)

	orderService := NewOrderService()
	orderService.O.f2(context.Background(), 1)
}
