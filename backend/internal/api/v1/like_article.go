package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/domain/dto"
	"github.com/joey17520/ailiaili/internal/resp"
	"github.com/joey17520/ailiaili/internal/service"
	"github.com/joey17520/ailiaili/utils"
)

func LikeArticle(ctx *gin.Context) {
	// 获取参数
	var likeReq dto.LikeArticleReq
	if err := ctx.Bind(&likeReq); err != nil {
		resp.FailWithMessage(ctx, "请求参数有误")
		return
	}

	if err := service.LikeArticle(ctx, likeReq); err != nil {
		resp.FailWithMessage(ctx, err.Error())
		return
	}

	// 返回
	resp.Ok(ctx)
}

func CancelLikeArticle(ctx *gin.Context) {
	var likeReq dto.LikeArticleReq
	if err := ctx.Bind(&likeReq); err != nil {
		resp.FailWithMessage(ctx, "请求参数有误")
		return
	}

	if err := service.CancelLikeArticle(ctx, likeReq); err != nil {
		resp.FailWithMessage(ctx, err.Error())
		return
	}

	// 返回
	resp.Ok(ctx)
}

func HasLikeArticle(ctx *gin.Context) {
	articleId := utils.StringToUint(ctx.Query("aid"))
	like, err := service.HasLikeArticle(ctx, articleId)
	if err != nil {
		resp.FailWithMessage(ctx, err.Error())
		return
	}

	// 返回
	resp.OkWithData(ctx, gin.H{"like": like})
}
