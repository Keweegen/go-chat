package cache

import (
    "context"
    "errors"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
)

const ClientRedis = iota

type Cache struct {
    native *redis.Client
}

type Connection struct {
    Client   int
    Host     string
    Port     int
    Database int
    Username string
    Password string
}

var cache *Cache

func GetCache(opts Connection) (*Cache, error) {
    if cache != nil {
        return cache, nil
    }

    var client *redis.Client
    switch opts.Client {
    case ClientRedis:
        client = opts.getRedisClient()
    default:
        return nil, errors.New("undefined cache client")
    }

    cache = &Cache{}
    cache.native = client

    if err := cache.Ping(context.Background()); err != nil {
        return nil, err
    }

    return cache, nil
}

func (opts Connection) getRedisClient() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", opts.Host, opts.Port),
        Username: opts.Username,
        Password: opts.Password,
        DB:       opts.Database,
    })
}

func (c *Cache) Ping(ctx context.Context) error {
    return c.native.Ping(ctx).Err()
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
    return c.native.Set(ctx, key, value, expiration).Err()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
    return c.native.Get(ctx, key).Result()
}

func (c *Cache) LPush(ctx context.Context, key string, values ...interface{}) error {
    return c.native.LPush(ctx, key, values...).Err()
}

func (c *Cache) RPush(ctx context.Context, key string, values ...interface{}) error {
    return c.native.RPush(ctx, key, values...).Err()
}

func (c *Cache) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
    return c.native.BLPop(ctx, timeout, keys...).Result()
}

func (c *Cache) BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
    return c.native.BRPop(ctx, timeout, keys...).Result()
}
