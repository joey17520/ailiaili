package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectDanmakuRoutes(r *gin.RouterGroup) {
	danmakuGroup := r.Group("danmaku")

	danmakuAuth := danmakuGroup.Group("")
	danmakuAuth.Use(middleware.Auth())
	{
		// 发送弹幕
		danmakuAuth.POST("sendDanmaku", api.SendDanmaku)
	}

	// 获取弹幕列表
	danmakuGroup.GET("getDanmaku", api.GetDanmaku)
}
