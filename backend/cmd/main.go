package main

import (
	"flag"

	"github.com/joey17520/ailiaili/internal/cron"
	"github.com/joey17520/ailiaili/internal/global"
	"github.com/joey17520/ailiaili/internal/initialize"
	"github.com/joey17520/ailiaili/internal/routes"
	"github.com/joey17520/ailiaili/internal/service"
	"github.com/joey17520/ailiaili/pkg/casbin"
	"github.com/joey17520/ailiaili/pkg/jigsaw"
	"github.com/joey17520/ailiaili/pkg/logger"
	"github.com/joey17520/ailiaili/pkg/mysql"
	"github.com/joey17520/ailiaili/pkg/oss"
	"github.com/joey17520/ailiaili/pkg/redis"
)

func main() {
	env := flag.String("env", "prod", "dev/prod")
	flag.Parse()

	// 初始化配置文件
	initialize.InitConfig(*env)
	// 初始化日志
	logger.InitLogger()
	// 初始化滑块验证码生成
	jigsaw.Jigsaw()
	// 初始化OSS
	if global.Config.Storage.OssType != "local" {
		global.Storage = oss.InitStorage(global.Config.Storage)
	}
	// 初始化雪花ID
	initialize.InitSnowflake()
	// 初始化mysql
	global.Mysql = mysql.Init(global.Config.Mysql)
	initialize.InitTables()
	initialize.InitDefaultData()
	// 初始化分区数据
	global.VideoPartitionMap = service.GetPartitionMap(global.CONTENT_TYPE_VIDEO)
	// 初始化缓存
	global.Redis = redis.Init(global.Config.Redis)
	initialize.InitCacheData()
	// 初始化casbin
	global.Casbin = casbin.InitCasbin()

	// 手动执行一次刷新热点视频
	cron.RefreshPopular()
	// 启动定时任务
	go cron.StartCronTask()

	// 初始化路由
	routes.InitRouter()
}
