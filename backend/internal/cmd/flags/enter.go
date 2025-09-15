package flags

import (
	"flag"
	"os"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
)

type Options struct {
	File    string
	DB      bool
	Version bool
}

var FlagOptions = new(Options)

func Parse() {
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.Parse()
}
func Run() {
	if FlagOptions.DB {
		//执行数据库迁移
		migrateTables()
		os.Exit(0)
	}
}
func migrateTables() {
	//自动迁移某一个表，确保表结构存在
	err := global.DB.AutoMigrate(
		&model.User{},
		&model.Order{},
		&model.Summary{},
	)
	if err != nil {
		global.Log.Error("数据库迁移失败")
	}
	global.Log.Info("数据库迁移成功！")
}
