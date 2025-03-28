package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/resp"
	"github.com/joey17520/ailiaili/internal/service"
	"github.com/joey17520/ailiaili/utils"
)

func GetLikeMessage(ctx *gin.Context) {
	page := utils.StringToInt(ctx.Query("page"))
	pageSize := utils.StringToInt(ctx.Query("pageSize"))

	if pageSize > 30 {
		resp.FailWithMessage(ctx, "请求数量过多")
		return
	}

	total, messages := service.GetLikeMessage(ctx, page, pageSize)

	// 返回给前端
	resp.OkWithData(ctx, gin.H{"total": total, "messages": messages})
}
