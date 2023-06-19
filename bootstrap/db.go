package bootstrap

import (
	"github.com/wecanooo/gosari/core"
	"github.com/wecanooo/gosari/database"
)

/**
 * 마이그레이트를 진행할 테이블을 넘겨줘야 함
 * ex,
 *
 */
func SetupDB(dst ...interface{}) {
	db, sqlDB := database.SetupDefaultDatabase()
	if core.GetConfig().Bool("DB.MASTER.AUTO_MIGRATE") {
		err := db.AutoMigrate(dst)
		if err != nil {
			return
		}
	}
	core.NewDatabaseConnection(db, sqlDB)
}

func SetupRedis() {
	client := database.SetupRedis()
	core.NewRedisConnection(client)
}
