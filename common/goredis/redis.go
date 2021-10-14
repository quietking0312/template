package goredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"runtime"
	"time"
)

var _redisClient *redis.Client

type redisCfg struct {
	Addr         string
	PoolSize     int
	MinIdleConns int
	Username     string
	Password     string
	DB           uint32
}

type Option func(cfg *redis.Options)

func Addr(ip string, port int) Option {
	return func(cfg *redis.Options) {
		cfg.Addr = fmt.Sprintf("%s:%d", ip, port)
	}
}

func PoolSize(poolSize int) Option {
	return func(cfg *redis.Options) {
		cfg.PoolSize = poolSize
	}
}

func MinIdleConns(idle int) Option {
	return func(cfg *redis.Options) {
		cfg.MinIdleConns = idle
	}
}

func Username(username string) Option {
	return func(cfg *redis.Options) {
		cfg.Username = username
	}
}

func Password(password string) Option {
	return func(cfg *redis.Options) {
		cfg.Password = password
	}
}

func DBName(db int) Option {
	return func(cfg *redis.Options) {
		cfg.DB = db
	}
}

func defaultRedisOption() *redis.Options {
	return &redis.Options{
		Network:         "tcp",
		Addr:            "127.0.0.1:6379",
		Dialer:          nil,
		OnConnect:       nil,
		Username:        "",
		Password:        "",
		DB:              0,
		MaxRetries:      3,                // 最大重试测试； -1 禁用重试
		MinRetryBackoff: time.Duration(8), // 重试直接的最小
		MaxRetryBackoff: time.Duration(512),
		// 新连接超时时间
		DialTimeout: 5 * time.Second,
		// 读取超时
		ReadTimeout: 3 * time.Second,
		//写入超时
		WriteTimeout: 3 * time.Second,
		// 连接池大小
		PoolSize: runtime.NumCPU() * 10,
		// 空闲连接数
		MinIdleConns: 5,
		// 连接过时 时间
		MaxConnAge: 0 * time.Hour,
		// 池超时
		PoolTimeout: 4 * time.Second,
		// 空闲超时时间
		IdleTimeout: 5 * time.Minute,
		// 空闲连接超时检测 频率
		IdleCheckFrequency: 1 * time.Minute,
		// TLS 设置
		TLSConfig: nil,
		// 限制器
		Limiter: nil,
	}
}

func NewRedisClient(ctx context.Context, opts ...Option) (*redis.Client, error) {
	redisCfg := defaultRedisOption()
	for _, opt := range opts {
		opt(redisCfg)
	}
	_redisClient = redis.NewClient(redisCfg)
	_, err := _redisClient.Ping(ctx).Result()
	return _redisClient, err
}

func GetConn(ctx context.Context) (*redis.Conn, error) {
	conn := _redisClient.Conn(ctx)
	_, err := conn.Ping(ctx).Result()
	return conn, err
}
