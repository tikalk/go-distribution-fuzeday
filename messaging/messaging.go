package messaging

import (
	"fmt"
	"github.com/matryer/vice"
	"github.com/matryer/vice/queues/redis"
	redisv3 "gopkg.in/redis.v3"
)

var transport vice.Transport
var provider = Redis

type Provider uint16

const (
	Redis    Provider = 0
	RabbitMQ Provider = 1
	ActiveMQ Provider = 2
	PubSub   Provider = 3
	SQS      Provider = 4
)

const BallChannelName = "ball_status"
const DisplayChannelName = "display"

const LocalAddr = "127.0.0.1:6379"
const LocalPass = ""

var RedisAddr = LocalAddr
var RedisPass = LocalPass

func getTransport() vice.Transport {
	if transport == nil {
		switch provider {
		case Redis:
			client := redisv3.NewClient(&redisv3.Options{
				Network:    "tcp",
				Addr:       RedisAddr,
				Password:   RedisPass,
				DB:         0,
				MaxRetries: 0,
			})
			transport = redis.New(redis.WithClient(client))
		case RabbitMQ:
		case ActiveMQ:
		case PubSub:
		case SQS:
			transport = nil
		}
	}
	return transport
}

func GetOutputChannel(name string) chan<- []byte {
	if transport == nil {
		getTransport()
	}
	fmt.Printf("GetOutputChannel: %s\n", name)
	return transport.Send(name)
}

func GetInputChannel(name string) <-chan []byte {
	if transport == nil {
		getTransport()
	}
	fmt.Printf("GetInputChannel: %s\n", name)
	return transport.Receive(name)
}
