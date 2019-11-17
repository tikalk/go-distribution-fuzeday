package commands

import (
	"fmt"
	"github.com/urfave/cli"
	"go-distribution-fuzeday/messaging"
)

func setupRedis(c *cli.Context) {
	redisHost := c.GlobalString("redis-host")
	redisPort := c.GlobalInt("redis-port")
	redisAddr := fmt.Sprintf("%s:%d", redisHost, redisPort)
	messaging.RedisAddr = redisAddr
}
