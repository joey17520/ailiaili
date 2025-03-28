package cache

import (
	"encoding/json"

	"github.com/joey17520/ailiaili/internal/domain/vo"
	"github.com/joey17520/ailiaili/internal/global"
	"github.com/joey17520/ailiaili/utils"
)

func GetUserInfo(id uint) (user vo.UserInfoResp) {
	jsonStr := global.Redis.Get(USER_INFO_KEY + utils.UintToString(id))
	if jsonStr == "" {
		return
	}

	// 反序列化
	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		utils.ErrorLog("用户信息反序列化失败", "cache", err.Error())
	}
	return
}

func SetUserInfo(user vo.UserInfoResp) {
	//先序列化user
	ub, err := json.Marshal(user)
	if err != nil {
		utils.ErrorLog("用户信息序列化失败", "cache", err.Error())
		return
	}

	global.Redis.Set(USER_INFO_KEY+utils.UintToString(user.ID), ub, USER_INFO_EXPIRATION_TIME)
}

func DelUserInfo(id uint) {
	global.Redis.Del(USER_INFO_KEY + utils.UintToString(id))
}
