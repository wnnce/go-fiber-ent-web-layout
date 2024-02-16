package data

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go-fiber-ent-web-layout/ent"
	"go-fiber-ent-web-layout/internal/conf"
	"time"
)

var InjectSet = wire.NewSet(NewData, NewExampleRepo)

type Data struct {
	Ec *ent.Client    // ent客户端
	Rc *RedisOptional // 封装的redis操作
}

type RedisOptional struct {
	rc *redis.Client
}

func (r *RedisOptional) Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) error {
	valueStr, err := sonic.Marshal(value)
	if err != nil {
		return err
	}
	_, err = r.rc.Set(ctx, key, valueStr, expireTime).Result()
	return err
}

func (r *RedisOptional) Get(ctx context.Context, key string, value interface{}) error {
	result, err := r.rc.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return sonic.UnmarshalString(result, value)
}

func (r *RedisOptional) Remove(ctx context.Context, key string) error {
	_, err := r.rc.Del(ctx, key).Result()
	return err
}

func NewData(conf *conf.Data) (*Data, func(), error) {
	client, err := ent.Open(conf.Database.Driver, fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.DbName, conf.Database.Password))
	if err != nil {
		return nil, nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		DB:           conf.Redis.Index,
		Username:     conf.Redis.Username,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout,
		WriteTimeout: conf.Redis.WireTimeout,
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		_ = client.Close()
		_ = rdb.Close()
	}
	return &Data{
		Ec: client,
		Rc: &RedisOptional{
			rc: rdb,
		},
	}, cleanup, nil
}
