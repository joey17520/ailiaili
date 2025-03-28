package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/service"
	"github.com/joey17520/ailiaili/utils"
)

// 视频Websocket连接(统计在线人数)
func GetVideoOnlineConnect(ctx *gin.Context) {
	videoId := utils.StringToUint(ctx.Query("vid"))
	clientId := ctx.Query("clientId")
	if videoId == 0 || clientId == "" {
		return
	}

	// 升级为websocket长链接
	service.GetVideoOnlineConnect(ctx, videoId, clientId)
}
