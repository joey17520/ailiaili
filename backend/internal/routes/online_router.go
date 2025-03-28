package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
)

func CollectOnlineRoutes(r *gin.RouterGroup) {
	onlineGroup := r.Group("online")

	// 视频Websocket连接(统计在线人数)
	onlineGroup.GET("video", api.GetVideoOnlineConnect)
}
