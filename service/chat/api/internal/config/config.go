package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	RedisConf        redis.RedisConf
	ChatRpcConf      zrpc.RpcClientConf
	UserRpcConf      zrpc.RpcClientConf
	ChatKqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
