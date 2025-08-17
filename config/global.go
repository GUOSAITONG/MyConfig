package config

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config Appconf
	Ctx    = context.Background()
	Rdb    *redis.Client
	ES     *elasticsearch.Client
)
