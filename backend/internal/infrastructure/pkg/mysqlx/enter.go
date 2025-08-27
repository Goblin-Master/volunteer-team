package mysqlx

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"volunteer-team/backend/internal/infrastructure/configs"
)

func InitMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(configs.Conf.DB.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //不生成外键约束
		Logger:                                   logger.Default.LogMode(logger.Warn),
		PrepareStmt:                              true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
