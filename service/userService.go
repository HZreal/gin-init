package service

import (
	"encoding/json"
	"fmt"
	"gin-init/model/dto"
	"gin-init/model/entity"
	"gin-init/model/vo"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type UserService struct {
	UserModel *entity.UserModel
	// RedisService *RedisService // TODO 封装
}

func NewUserService(userModel *entity.UserModel) *UserService {
	return &UserService{UserModel: userModel}
}

func (uS *UserService) GetUserList(c *gin.Context, query dto.QueryPagination, body map[string]interface{}) (result *vo.PaginationResult) {
	//

	var userInfos []vo.UserSingle
	var total int64
	page, pageSize := query.Page, query.PageSize

	//
	offset := (page - 1) * pageSize

	// 获取数据总数和分页数据
	// db.Model(&entity.UserModel{}).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)
	// TODO 通过依赖注入
	db.Model(uS.UserModel).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)

	// 计算总页数
	pages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		pages++
	}

	//
	return &vo.PaginationResult{
		Total:       int(total),
		Pages:       pages,
		CurrentPage: page,
		PageSize:    pageSize,
		Records:     userInfos,
	}
}

func (uS *UserService) GetUserDetail(c *gin.Context, id string) (userInfo vo.UserSingle) {
	//
	key := fmt.Sprintf("tmp:user:id:%s", id)
	cachedData, err := rdb.Get(c, key).Result()
	if err == redis.Nil {
		// 无缓存
		affected := db.Take(&entity.UserModel{}, id).Scan(&userInfo).RowsAffected
		if affected == 0 {
			log.Printf("No user found with ID: %s", id)
			return
		}

		// 将查询结果序列化为 JSON 字符串
		unitInfoJson, err := json.Marshal(userInfo)
		if err != nil {
			log.Printf("Failed to serialize data for user ID: %s, error: %v", id, err)
			panic("failed to serialize data")
		}

		// 将数据缓存到 Redis，设置缓存过期时间为 30 S
		err = rdb.Set(c, key, unitInfoJson, 30*time.Second).Err()
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
			log.Printf("Failed to deserialize cache data for user ID: %s, error: %v", id, err)
			panic("failed to deserialize cache data")
		}

		return userInfo
	}

}

func (uS *UserService) CheckUser(loginData dto.LoginData) bool {
	//
	username, password := loginData.Username, loginData.Password

	// TODO（这里应该是从数据库中验证用户信息）
	if !(username == "admin" && password == "root123456") {
		return false
	}
	return true
}
