package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/domain/dto"
	"github.com/joey17520/ailiaili/internal/domain/model"
	"github.com/joey17520/ailiaili/internal/domain/vo"
	"github.com/joey17520/ailiaili/internal/global"
	"github.com/joey17520/ailiaili/utils"
)

func GetDanmaku(ctx *gin.Context, videoId, part uint) (danmakus []vo.DanmakuResp, err error) {
	// 查询用户信息
	if err := global.Mysql.Model(&model.Danmaku{}).Select(vo.DANMAKU_FIELD).
		Where("vid = ? and part = ?", videoId, part).Scan(&danmakus).Error; err != nil {
		utils.ErrorLog("获取粉丝列表失败", "relation", err.Error())
		return danmakus, errors.New("获取失败")
	}

	return danmakus, nil
}

func SendDanmaku(ctx *gin.Context, danmakuReq dto.DanmakuReq) error {
	if video := GetVideoInfo(danmakuReq.Vid); video.ID == 0 {
		return errors.New("视频不存在")
	}

	userId := ctx.GetUint("userId")
	if danmakuReq.Part == 0 {
		danmakuReq.Part = 1
	}

	// 保存到数据库
	if err := global.Mysql.Create(&model.Danmaku{
		Vid:   danmakuReq.Vid,
		Uid:   userId,
		Part:  danmakuReq.Part,
		Time:  danmakuReq.Time,
		Type:  danmakuReq.Type,
		Color: danmakuReq.Color,
		Text:  danmakuReq.Text,
	}).Error; err != nil {
		utils.ErrorLog("保存弹幕失败", "danmaku", err.Error())
		return errors.New("发送失败")
	}

	return nil
}
