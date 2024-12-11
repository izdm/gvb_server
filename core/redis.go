package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password, // no password set
		DB:       db,                 // use default DB
		PoolSize: redisConf.PoolSize, // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cmd, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("redis连接失败 %s: %v", redisConf.Addr(), err)
		return nil
	}
	logrus.Infof("redis连接成功，返回结果：%s", cmd)
	return rdb
}
