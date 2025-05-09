package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/domain/dto"
	"github.com/joey17520/ailiaili/internal/resp"
	"github.com/joey17520/ailiaili/internal/service"
	"github.com/joey17520/ailiaili/utils"
)

func CollectArticle(ctx *gin.Context) {
	// 获取参数
	var collectReq dto.CollectArticleReq
	if err := ctx.Bind(&collectReq); err != nil {
		resp.FailWithMessage(ctx, "请求参数有误")
		return
	}

	if err := service.CollectArticle(ctx, collectReq); err != nil {
		resp.FailWithMessage(ctx, err.Error())
		return
	}

	// 返回
	resp.Ok(ctx)
}

func CancelCollectArticle(ctx *gin.Context) {
	var collectReq dto.CollectArticleReq
	if err := ctx.Bind(&collectReq); err != nil {
		resp.FailWithMessage(ctx, "请求参数有误")
		return
	}

	if err := service.CancelCollectArticle(ctx, collectReq); err != nil {
		resp.FailWithMessage(ctx, err.Error())
		return
	}

	// 返回
	resp.Ok(ctx)
}

func HasCollectArticle(ctx *gin.Context) {
	articleId := utils.StringToUint(ctx.Query("aid"))
	collect, err := service.HasCollectArticle(ctx, articleId)
	if err != nil {
		resp.FailWithMessage(ctx, err.Error())
		return
	}

	// 返回
	resp.OkWithData(ctx, gin.H{"collect": collect})
}
