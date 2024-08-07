package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-init/model/dto"
	"gin-init/model/entity"
	"gin-init/model/vo"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

type UserService struct {
	UserModel *entity.UserModel
}

func (uS *UserService) GetUserList(c *gin.Context, query dto.QueryPagination, body map[string]interface{}) (result vo.PaginationResult, err error) {
	//

	var userInfos []vo.UserSingle
	var total int64
	page, pageSize := query.Page, query.PageSize

	//
	offset := (page - 1) * pageSize

	// 获取数据总数和分页数据
	db.Model(&entity.UserModel{}).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)
	// TODO error
	// db.Model(uS.UserModel).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)

	// 计算总页数
	pages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		pages++
	}

	//
	return vo.PaginationResult{
		Total:       int(total),
		Pages:       pages,
		CurrentPage: page,
		PageSize:    pageSize,
		Records:     userInfos,
	}, nil
}

func (uS *UserService) GetUserDetail(c *gin.Context, id string) (userInfo vo.UserSingle, err error) {

	//
	//
	key := fmt.Sprintf("tmp:user:id:%s", id)
	cachedData, err := rdb.Get(c, key).Result()
	if err == redis.Nil {
		// 无缓存
		affected := db.Take(&entity.UserModel{}, id).Scan(&userInfo).RowsAffected
		if affected == 0 {
			// c.JSON(http.StatusInternalServerError, gin.H{"message": "query error"})
			return userInfo, errors.New("query error")
		}

		// 将查询结果序列化为 JSON 字符串
		unitInfoJson, err := json.Marshal(userInfo)
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to serialize data"})
			return userInfo, errors.New("failed to serialize data")
		}

		// 将数据缓存到 Redis，设置缓存过期时间为 30 S
		err = rdb.Set(c, key, unitInfoJson, 30*time.Second).Err()
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to set cache"})
			return userInfo, errors.New("failed to set cache")
		}

		// c.JSON(http.StatusOK, userInfo)
		return userInfo, nil

	} else if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get cache"})
		return userInfo, errors.New("failed to get cache")
	} else {
		// 如果缓存中有数据，返回缓存数据
		if err := json.Unmarshal([]byte(cachedData), &userInfo); err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to deserialize cache data"})
			return userInfo, errors.New("failed to deserialize cache data")
		}

		// c.JSON(http.StatusOK, userInfo)
		return userInfo, nil
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
