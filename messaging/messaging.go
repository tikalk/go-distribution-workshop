package messaging

import (
	"fmt"
)


const BallChannelName = "ball_status"

const defaultChannelBuffer = 1024

const LocalAddr = "127.0.0.1:6379"
const LocalPass = ""

var RedisAddr = LocalAddr
var RedisPass = LocalPass


// TODO Challenge get a package-internal transport object here

var channels map[string]chan []byte

func init(){
	channels = make(map[string]chan []byte)
}

func getChannel(name string) chan []byte {
	if _, ok := channels[name]; !ok {
		fmt.Printf("creating channel %s with size %d\n", name, defaultChannelBuffer)
		channels[name] = make(chan []byte, defaultChannelBuffer)
	}
	return channels[name]
}


// Used mainly to limit access direction on channel
func GetOutputChannel(name string) chan<- []byte {

	fmt.Printf("GetOutputChannel: %s\n", name)

	// TODO Challenge: Get a Publisher channel here via transport.Send
	return getChannel(name)

}

// Used mainly to limit access direction on channel
func GetInputChannel(name string) <-chan []byte {

	fmt.Printf("GetInputChannel: %s\n", name)

	// TODO Challenge: Get a Consumer channel here via transport.Send
	return getChannel(name)
}





