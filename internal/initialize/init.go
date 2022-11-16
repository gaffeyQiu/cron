package initialize

import (
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	ES  *elasticsearch.Client
	RDB *redis.Client
)

func InitGlobal() {
	DB = GormDB()
	ES = GetES()
	RDB = GetRedis()
}
