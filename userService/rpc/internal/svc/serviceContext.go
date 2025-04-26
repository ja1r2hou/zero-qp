package svc

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"zero-qp/userService/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Client //单机
	Mongo  *mongo.Database
}

func NewServiceContext(c config.Config) *ServiceContext {
	//mongo
	clientOptions := options.Client().ApplyURI(c.Mongo.Url)
	clientOptions.SetAuth(options.Credential{
		Username: c.Mongo.UserName,
		Password: c.Mongo.Password,
	})
	clientOptions.SetMinPoolSize(uint64(c.Mongo.MinPoolSize))
	clientOptions.SetMaxPoolSize(uint64(c.Mongo.MaxPoolSize))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		Redis: redis.NewClient(&redis.Options{
			Addr:         c.RedisCof.Addr,
			PoolSize:     c.RedisCof.PoolSize,
			MinIdleConns: c.RedisCof.MinIdleConns,
			Password:     c.RedisCof.Password,
		}),
		Mongo: client.Database(c.Mongo.Db),
	}
}
