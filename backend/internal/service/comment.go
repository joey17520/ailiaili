package service

import (
	"github.com/joey17520/ailiaili/internal/domain/model"
	"github.com/joey17520/ailiaili/internal/global"
)

func FindCommentById(id uint) (comment model.Comment, err error) {
	err = global.Mysql.First(&comment, id).Error
	return
}
