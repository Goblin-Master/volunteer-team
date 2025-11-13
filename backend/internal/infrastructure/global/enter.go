package global

import (
	"volunteer-team/backend/internal/infrastructure/pkg/mysqlx"
	"volunteer-team/backend/internal/infrastructure/pkg/syncx"
	"volunteer-team/backend/internal/infrastructure/pkg/zapx"
	"volunteer-team/backend/internal/infrastructure/utils/snowflake"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	RDB       redis.Cmdable
	Log       *zap.SugaredLogger
	CodeStore *syncx.Map[string, string]
	Node      *snowflake.Node
)

func Init() {
	Log = zapx.InitZap()
	DB = mysqlx.InitMysql()
	//RDB = redisx.InitRedis()
	CodeStore = new(syncx.Map[string, string])
	Node, _ = snowflake.NewNode(DEFAULT_NODE_ID)
}
