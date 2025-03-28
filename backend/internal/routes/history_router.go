package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectHistoryRoutes(r *gin.RouterGroup) {
	historyGroup := r.Group("history")

	historyAuth := historyGroup.Group("")
	historyAuth.Use(middleware.Auth())
	{
		// 记录历史记录
		historyAuth.POST("video/addHistory", api.AddHistory)
		// 获取历史记录
		historyAuth.GET("video/getHistory", api.GetHistoryList)
		// 获取播放进度
		historyAuth.GET("video/getProgress", api.GetHistoryProgress)
	}
}
