package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/joey17520/ailiaili/internal/global"
	"github.com/joey17520/ailiaili/utils"
)

func InitSnowflake() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		utils.ErrorLog("雪花ID初始化失败", "snowflake", err.Error())
	}

	global.SnowflakeNode = node
}
