package commands

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/tikalk/go-distribution-workshop/messaging"
)

func setupRedis(c *cli.Context){
	redisHost := c.GlobalString("redis-host")
	redisPort := c.GlobalInt("redis-port")
	redisAddr := fmt.Sprintf("%s:%d", redisHost, redisPort)
	messaging.RedisAddr = redisAddr
}
