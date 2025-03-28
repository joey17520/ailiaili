package mysql

import (
	"fmt"

	"github.com/joey17520/ailiaili/internal/config"
	"github.com/joey17520/ailiaili/utils"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var db *gorm.DB

func Init(c config.Mysql) *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Datasource, c.Param)
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()
	if mysqlClient, err := gorm.Open(mysql.Open(dns), &gorm.Config{Logger: logger}); err != nil {
		utils.ErrorLog("mysql连接失败", "db", err.Error())
		panic(err)
	} else {
		zap.L().Info("mysql连接成功", zap.String("module", "db"))
		db = mysqlClient
		return mysqlClient
	}
}

func GetMysqlClient() *gorm.DB {
	return db
}
