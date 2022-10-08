package redis

import (
	"context"
	"fmt"
	"github.com/crazy-choose/helper/log"
	"github.com/crazy-choose/helper/model/meta"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	rc         *redis.Client // rds 连接
	cfg        meta.Redis
	ctx        context.Context // rds 上下文
	expireTime = 86400 * time.Second
)

func Init(opt meta.Redis) bool {
	cfg = opt
	rc = redis.NewClient(&redis.Options{
		Addr:         cfg.Host,
		Password:     cfg.AuthKey,  // no password set
		DB:           cfg.DbIndex,  // use default DB
		PoolSize:     cfg.PoolSize, // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: 10,           //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；
		//超时
		DialTimeout:  cfg.DialTimeout * time.Second,  //连接建立超时时间，默认5秒。
		ReadTimeout:  cfg.ReadTimeout * time.Second,  //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: cfg.WriteTimeout * time.Second, //写超时，默认等于读超时
		PoolTimeout:  cfg.WriteTimeout * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒
	})
	ctx = context.Background()
	pong, err := rc.Ping(ctx).Result()
	if err != nil {
		log.Error(fmt.Sprintf("rdb.Pong(%+v), err%+v", pong, err))
		return false
	}

	log.Info("InitRedis Success")
	return true
}

func Impl() *redis.Client {
	if rc == nil {
		Init(cfg)
	}
	return rc
}

func Ctx() context.Context {
	return ctx
}

func Get(key string) *redis.StringCmd {
	return rc.Get(ctx, key)
}

func DoGet(key string) (interface{}, error) {
	return rc.Do(ctx, "GET", key).Result()
}

func Set(key string, value interface{}, expire time.Duration) *redis.StatusCmd {
	return rc.Set(ctx, key, value, expire)
}

func Publish(channel string, message interface{}) *redis.IntCmd {
	return rc.Publish(ctx, channel, message)
}

func Subscribe(channel ...string) *redis.PubSub {
	return rc.Subscribe(ctx, channel...)
}

func PSubscribe(handler func(ps *redis.PubSub), channel ...string) error {
	handler(rc.PSubscribe(ctx, channel...))
	return nil
}

func SMembers(key string) *redis.StringSliceCmd {
	return rc.SMembers(ctx, key)
}

func ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return rc.ZRange(ctx, key, start, stop)
}

func HGetAll(key string) *redis.StringStringMapCmd {
	return rc.HGetAll(ctx, key)
}

func LRange(key string, start, stop int64) *redis.StringSliceCmd {
	return rc.LRange(ctx, key, start, stop)
}

func LPush(key string, val interface{}) *redis.IntCmd {
	return rc.LPush(ctx, key, val)
}

func DelKey(key string) *redis.IntCmd {
	return rc.Del(ctx, key)
}

func SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	//	args := []interface{}{"SSCAN", key, cursor, match, count}
	return rc.SScan(ctx, key, cursor, match, count)
}

func ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return rc.ZScan(ctx, key, cursor, match, count)
}

func Flush() {
	rc.FlushDB(ctx) // 清除数据
}
