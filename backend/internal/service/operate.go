package service

import (
	"github.com/joey17520/ailiaili/internal/domain/model"
	"github.com/joey17520/ailiaili/internal/global"
)

func AddOperate(operate *model.Operate) error {
	return global.Mysql.Create(&operate).Error
}
