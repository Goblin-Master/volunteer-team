package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"volunteer-team/backend/internal/infrastructure/pkg/mysqlx"
	"volunteer-team/backend/internal/infrastructure/pkg/syncx"
	"volunteer-team/backend/internal/infrastructure/pkg/zapx"
)

var (
	DB        *gorm.DB
	RDB       redis.Cmdable
	Log       *zap.SugaredLogger
	CodeStore *syncx.Map[string, string]
)

func Init() {
	Log = zapx.InitZap()
	DB = mysqlx.InitMysql()
	//RDB = redisx.InitRedis()
	CodeStore = new(syncx.Map[string, string])
}
