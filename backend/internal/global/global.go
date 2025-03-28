package global

import (
	"github.com/bwmarrin/snowflake"
	"github.com/joey17520/ailiaili/internal/config"
	"github.com/joey17520/ailiaili/pkg/casbin"
	"github.com/joey17520/ailiaili/pkg/oss"
	"github.com/joey17520/ailiaili/pkg/redis"
	"gorm.io/gorm"
)

var (
	Config            *config.Config
	Mysql             *gorm.DB
	Redis             *redis.Redis
	Casbin            *casbin.Casbin
	Storage           oss.Storage
	SnowflakeNode     *snowflake.Node
	VideoPartitionMap map[uint]uint
)
