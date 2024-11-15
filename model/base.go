package main

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

type BaseModelInterface interface {
	Insert(ctx context.Context, data *EntityType) (EntityType, error)
	FindOne(ctx context.Context, id int64) (EntityType, error)
	Update(ctx context.Context, data *EntityType) error
	Delete(ctx context.Context, id int64) error
}

type BaseRepo struct {
	conn  string
	table string
}

func NewBaseRepo() *BaseRepo {
	return &BaseRepo{conn: "", table: ""}
}

func (m *BaseRepo) Insert(ctx context.Context, data *EntityType) (EntityType, error) {
	return &UserEntity{Id: 1, Name: ""}, nil
}

func (m *BaseRepo) FindOne(ctx context.Context, id int64) (EntityType, error) {
	return &UserEntity{Id: 1, Name: ""}, nil
}

func (m *BaseRepo) Update(ctx context.Context, data *EntityType) error {
	return nil
}

func (m *BaseRepo) Delete(ctx context.Context, id int64) error {
	return nil
}

// ///////////////////////////////////////////////////////////////////////////////

type UserRepoInterface interface {
	BaseModelInterface

	f1(ctx context.Context, id int) bool
}

type UserEntity struct {
	Id   int
	Name string
}

type UserRepo struct {
	*BaseRepo
}

func (m *UserRepo) f1(ctx context.Context, id int) bool {
	return true
}

func NewUserRepo() UserRepoInterface {
	return &UserRepo{
		BaseRepo: NewBaseRepo(),
	}
}

// ///////////////////////////////////////////////////////////////////////////////

type OrderEntity struct {
	Id   int
	DESC string
}

type OrderRepoInterface interface {
	BaseModelInterface

	f2(id int) bool
}

type OrderRepo struct {
	*BaseRepo
}

func (m *OrderRepo) f2(id int) bool {
	return true
}

func NewOrderRepo() OrderRepoInterface {
	return &OrderRepo{
		BaseRepo: NewBaseRepo(),
	}
}

// ///////////////////////////////////////////////////////////////////////////////

type UserService struct {
	U UserRepoInterface
}

// ///////////////////////////////////////////////////////////////////////////////

type OrderService struct {
	O OrderRepoInterface
}

func main() {
	userService := &UserService{
		U: NewUserRepo(),
	}
	userService.U.FindOne(context.Background(), 1)

	orderService := &OrderService{
		O: NewOrderRepo(),
	}
	orderService.O.f2(1)
}
