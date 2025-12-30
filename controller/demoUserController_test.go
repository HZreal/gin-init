package controller

import (
	"bytes"
	"encoding/json"
	"gin-init/model/types"
	"gin-init/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// 注册用户相关接口
	r.POST("/api/user/create", func(c *gin.Context) {
		uc := NewUserController()
		uc.Create(c)
	})
	r.POST("/api/user/list", func(c *gin.Context) {
		uc := NewDemoUserController(&service.DemoUserService{})
		uc.GetAllUser(c)
	})
	r.GET("/api/user/detail", func(c *gin.Context) {
		uc := NewDemoUserController(&service.DemoUserService{})
		uc.GetUserDetail(c)
	})
	r.POST("/api/user/update", func(c *gin.Context) {
		uc := NewDemoUserController(&service.DemoUserService{})
		uc.UpdateUser(c)
	})
	r.POST("/api/user/delete", func(c *gin.Context) {
		uc := NewDemoUserController(&service.DemoUserService{})
		uc.DeleteUser(c)
	})
	return r
}

func TestCreateUser(t *testing.T) {
	r := setupRouter()
	body := types.UserCreateDTO{
		Username: "testuser",
		Password: "123456",
		Phone:    "13800138000",
		Age:      20,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/user/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetAllUser(t *testing.T) {
	r := setupRouter()
	body := types.UsersFilterDTO{
		Username: "testuser",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/user/list", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetUserDetail(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/api/user/detail?id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestUpdateUser(t *testing.T) {
	r := setupRouter()
	body := types.UserUpdateDTO{
		Id:       1,
		Username: "updateduser",
		Phone:    "13800138001",
		Age:      22,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/user/update", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	r := setupRouter()
	body := types.BodyJsonId{
		Id: 1,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/user/delete", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}
