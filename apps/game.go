package main

import (
	"time"
	"encoding/json"
	"math/rand"
	"sync"
	"fmt"
	"github.com/tikalk/go-distribution-workshop/messaging"
	"github.com/tikalk/go-distribution-workshop/models"
)

// TODO have all players and redis connectio parameters passed via CLI flags so that people with no Golang IDE can still participate in the workshop

func main() {
	defer messaging.Stop()

	throwBall()

	var numPlayers = 6

	displayChannel := getDisplayOutputChannel()
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	wg.Add(numPlayers)


	for i := 0; i < numPlayers; i++ {

		player := &models.Player{
			ID: fmt.Sprintf("Player %d", i),
			X: rand.Float64() * 100,
			Y: rand.Float64() * 100,
			MaxVelocity: rand.Float64() * 0.1,
		}
		if i %2 == 0 {
			player.TeamID = models.TeamRed
		} else {
			player.TeamID = models.TeamBlue
		}

		player.Activate(displayChannel, wg)
	}

	wg.Wait()


}

func getDisplayOutputChannel() chan <- *models.DisplayStatus  {
	rawOutput, _ := messaging.GetOutputChannel(messaging.DisplayChannelName)
	res := make(chan *models.DisplayStatus)

	// Display channel population, executed in function closure
	go func(){
		for bs := range res {
			val, err := json.Marshal(bs)
			if err == nil {
				rawOutput <- val
			}
		}
	}()
	return res
}

func throwBall(){
	output, _ := messaging.GetOutputChannel(messaging.BallChannelName)

	bs := &models.Ball{X: 8, Y: 2, Vx: 0, Vy: 0, Vz: 0, Z: 50}
	bs.LastUpdated = time.Now()
	bsSer, err := json.Marshal(bs)
	if err != nil {
		panic(err)
	}
	output <- bsSer
}
