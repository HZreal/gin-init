package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-init/common/constant"
	"gin-init/core/initialize/database"
	"gin-init/model"
	"gin-init/model/entity"
	"gin-init/model/types"
	"gin-init/service/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var _ UserServiceInterface = (*UserService2)(nil)

type UserServiceInterface interface {
	BaseServiceInterface[entity.TbUser]
	ChangePassword(id uint, oldPassword, newPassword string) error
}

type UserService2 struct {
	*BaseService[entity.TbUser]
	TbUserModel  model.TbUserModelInterface
	RedisService *common.RedisService
}

func NewUserService2() UserServiceInterface {
	return &UserService2{
		RedisService: common.NewRedisService(),
		TbUserModel:  model.NewTbUserModel(),
		BaseService:  NewBaseService[entity.TbUser](),
	}
}

func (s *UserService2) ChangePassword(id uint, oldPassword, newPassword string) error {
	user, err := s.TbUserModel.FindByID(id)
	if err != nil {
		return err
	}
	if user.Password != oldPassword {
		return errors.New("old password is incorrect")
	}
	user.Password = newPassword
	return s.TbUserModel.Update(user)
}

type UserService struct {
	UserModel    *entity.TbUser
	RedisService *common.RedisService
}

func NewUserService(userModel *entity.TbUser, redisService *common.RedisService) *UserService {
	return &UserService{
		UserModel:    userModel,
		RedisService: redisService,
	}
}

func (uS *UserService) GetAllUser(c *gin.Context, body types.UsersFilterDTO) []types.UserDetailInfo {
	var users []types.UserDetailInfo
	if err := database.DB.Model(uS.UserModel).Where(body).Find(&users).Error; err != nil {
		log.Printf("query users err:%v", err)
		panic(err)
	}
	return users
}

func (uS *UserService) GetUserList(c *gin.Context, query types.QueryPagination, body map[string]interface{}) (result *types.PaginationResult) {
	//

	var userInfos []types.UserDetailInfo
	var total int64
	page, pageSize := query.Page, query.PageSize

	//
	offset := (page - 1) * pageSize

	// 获取数据总数和分页数据
	// db.Model(&entity.TbUser{}).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)
	// TODO 通过依赖注入
	database.DB.Model(uS.UserModel).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)

	// 计算总页数
	pages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		pages++
	}

	//
	return &types.PaginationResult{
		Total:       int(total),
		Pages:       pages,
		CurrentPage: page,
		PageSize:    pageSize,
		Records:     userInfos,
	}
}

func (uS *UserService) GetUserDetail(c *gin.Context, id uint) (userInfo types.UserDetailInfo) {
	//
	key := fmt.Sprintf(constant.UserDetail, id)
	cachedData, err := uS.RedisService.Client.Get(c, key).Result()
	if errors.Is(err, redis.Nil) {
		// 无缓存
		affected := database.DB.Take(&entity.TbUser{}, id).Scan(&userInfo).RowsAffected
		if affected == 0 {
			log.Printf("No user found with ID: %d", id)
			return
		}

		// 将查询结果序列化为 JSON 字符串
		unitInfoJson, err := json.Marshal(userInfo)
		if err != nil {
			log.Printf("Failed to serialize data for user ID: %d, error: %v", id, err)
			panic("failed to serialize data")
		}

		// 将数据缓存到 Redis，设置缓存过期时间为 30 S
		err = uS.RedisService.Client.Set(c, key, unitInfoJson, 30*time.Second).Err()
		if err != nil {
			panic("failed to save data")
		}

		return userInfo

	} else if err != nil {
		log.Printf("Failed to get cache for key: %s, error: %v", key, err)
		panic("failed to get cache")
	} else {
		// 如果缓存中有数据，返回缓存数据
		if err := json.Unmarshal([]byte(cachedData), &userInfo); err != nil {
			log.Printf("Failed to deserialize cache data for user ID: %d, error: %v", id, err)
			panic("failed to deserialize cache data")
		}

		return userInfo
	}

}

func (uS *UserService) CheckUser(loginData types.LoginData) bool {
	//
	username, password := loginData.Username, loginData.Password

	// TODO（这里应该是从数据库中验证用户信息）
	if !(username == "admin" && password == "root123456") {
		return false
	}
	return true
}

func (uS *UserService) CreateUser(c *gin.Context, body types.UserCreateDTO) types.UserDetailInfo {
	user := entity.TbUser{
		Username: body.Username,
		Password: body.Password,
		Phone:    body.Phone,
		Age:      body.Age,
	}

	if result := database.DB.Create(&user); result.Error != nil {
		log.Printf("Failed to create user, error: %v", result.Error)
		panic("failed to create user")
	}
	return types.UserDetailInfo{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Age:      user.Age,
	}

}

func (uS *UserService) UpdateUser(c *gin.Context, body types.UserUpdateDTO) types.UserDetailInfo {
	id := body.Id
	var user entity.TbUser
	if result := database.DB.First(&user, id); result.Error != nil {
		log.Printf("Failed to find user, error: %v", result.Error)
		panic("failed to find user")
	}

	//
	result := database.DB.Model(&user).Where("id = ?", id).Updates(body)
	if result.Error != nil {
		log.Printf("Failed to update user, error: %v", result.Error)
		panic("failed to update user")
	}
	return types.UserDetailInfo{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Age:      user.Age,
	}
}

func (uS *UserService) DeleteUser(c *gin.Context, id int) {
	if result := database.DB.First(uS.UserModel, id); result.Error != nil {
		log.Printf("Failed to find user, error: %v", result.Error)
		panic("failed to find user")
	}

	//
	result := database.DB.Delete(uS.UserModel, id)
	if result.Error != nil {
		log.Printf("Failed to delete user, error: %v", result.Error)
		panic("failed to delete user")
	}

}
