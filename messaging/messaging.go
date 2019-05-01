package messaging

import (
	"github.com/matryer/vice/queues/redis"
	goredis "github.com/go-redis/redis"
	"github.com/matryer/vice"
	"fmt"
)

var transport vice.Transport
var provider = Redis

type Provider uint16

const (
	Redis Provider = 0
	Rabbit Provider = 1
)

const BallChannelName = "ball_status"
const DisplayChannelName = "display"

const RedisDB = "go-workshop"

const LocalAddr = "127.0.0.1:6379"
const RemoteAddr = "redis-19098.c55.eu-central-1-1.ec2.cloud.redislabs.com:19098"
const LocalPass = ""
const RemotePass = "q1w2e3r4"

const RedisAddr = RemoteAddr
const RedisPass= RemotePass


func init(){

	switch provider {
	case Redis:
		client := goredis.NewClient(&goredis.Options{
			Network:    "tcp",
			Addr:       RedisAddr,
			Password:   RedisPass,
			DB:         0,
			MaxRetries: 0,
		})
		transport = redis.New(redis.WithClient(client))
	case Rabbit:
		transport = nil
	}

	println()
}

func Stop(){
	if transport != nil {
		transport.Stop()
		<-transport.Done()
	}
}
func GetErrorChannel() <-chan error {
	return transport.ErrChan()
}
func GetOutputChannel(name string) (chan<- []byte, vice.Transport){
	fmt.Println("GetOutputChannel")
	return transport.Send(name), transport
}

func GetInputChannel(name string) (<-chan []byte, vice.Transport){
	fmt.Println("GetInputChannel")
	return transport.Receive(name), transport
}



