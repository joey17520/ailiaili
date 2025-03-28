package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/api/v1"
	"github.com/joey17520/ailiaili/internal/middleware"
)

func CollectUploadRoutes(r *gin.RouterGroup) {

	uploadGroup := r.Group("upload")
	uploadGroup.Use(middleware.Auth())
	{
		uploadGroup.POST("image", api.UploadImg)
		uploadGroup.POST("video/:vid", api.UploadVideoAdd)
		uploadGroup.POST("video", api.UploadVideoCreate)
		uploadGroup.POST("checkVideo", api.UploadVideoCheck)
		uploadGroup.POST("chunkVideo", api.UploadVideoChunk) // 分片上传视频
		uploadGroup.POST("mergeVideo", api.UploadVideoMerge) // 合并视频分片
	}
}
