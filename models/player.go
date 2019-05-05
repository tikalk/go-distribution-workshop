package models

import (
	"encoding/json"
	"fmt"
	"time"
	"math"
	"math/rand"
	"sync"
	"github.com/tikalk/go-distribution-workshop/messaging"
	"github.com/tikalk/go-distribution-workshop/utils"
)

const kickThreshold = 6
const kickVelocityThreshold = 4

type Player struct {
	X float64
	Y float64
	TeamID Team
	ID string
	Name string

	MaxVelocity float64
	LastKick time.Time

	ballInput <-chan *Ball
	ballOutput chan<- *Ball

	idleV     	float64
	idleVx    	float64
	idleVy   	 float64
	idleAngle 	float64

}

func (p *Player) Activate(wg sync.WaitGroup) {

	p.ballInput = getBallInputChannel()
	p.ballOutput = getBallOutputChannel()

	var ball *Ball


	go func() {
		nextDelay := 0 * time.Second
		for {
			select {
			case <-time.After(nextDelay):

				p.idleV = 0.5 + 0.5 * rand.Float64()
				p.idleAngle = math.Pi * 2 * rand.Float64()
				p.idleVx = math.Cos(p.idleAngle) * p.idleV
				p.idleVy = math.Sin(p.idleAngle) * p.idleV
				nextDelay = time.Duration(5.0 + rand.Float64() * 6.0) * time.Second
			}
		}
	}()

	// Closing distance to ball
	go func() {
		for {
			select {
			case <-time.After(200 * time.Millisecond):
				p.runToBall(ball)
			}
		}
	}()

	ticker := time.NewTicker(10 * time.Second)

	go func() {

		for {
			select {

				case ball = <-p.ballInput:
					ticker.Stop()
					distance := p.getDistanceToBall(ball)

					if distance < kickThreshold &&
						ball.GetSurfaceVelocity() < kickVelocityThreshold &&
						time.Now().Sub(ball.LastKick) > 1*time.Second {

						p.applyKick(ball)

					} else {

						time.Sleep(20 * time.Millisecond)
						ball.ApplyKinematics()

					}

					p.log(fmt.Sprintf("Current Position: (%f, %f), Ball Position: (%f, %f)", p.X, p.Y, ball.X, ball.Y))
					ball.LastUpdated = time.Now()

					p.ballOutput <- ball

				case <-ticker.C:						// Initial delay before game starts
					if ball == nil {
						p.log("Waiting for the ball...\n")
					}

			}
		}
	}()

	wg.Done()

}

func (p *Player) getDistanceToBall(ball *Ball) float64 {
	return math.Sqrt(math.Pow(p.X-ball.X, 2) + math.Pow(p.Y-ball.Y, 2))
}

func (p *Player) runToBall(ball *Ball){

	// TODO make view threshold (50) random so that a distant player sees that ball after a period of time

	// once every N seconds - the player gets a longer view and can see the ball. Once saw the ball -
	// he keeps the "long view" mode for a longer period

	if ball != nil {
		dist := p.getDistanceToBall(ball)
		if dist < 50 && time.Now().Sub(p.LastKick) > 2 * time.Second{
			vel := 0.05 + rand.Float64() * p.MaxVelocity
			p.X += (ball.X - p.X) * vel
			p.Y += (ball.Y - p.Y) * vel
		} else {
			utils.ApplyVelocityComponent(&p.X, &p.idleVx, 1, 1)
			utils.ApplyVelocityComponent(&p.Y, &p.idleVy, 1, 1)
		}
	}

	p.log(fmt.Sprintf("Current Position: (%f, %f), Ball Position: (%f, %f)", p.X, p.Y, ball.X, ball.Y))


}

func (p *Player) log(message string) {
	if message[0:1] == "\n" {
		message = message[1:]
		fmt.Printf("\n%s: %s", p.Name, message)
	} else {
		fmt.Printf("\r%s: %s", p.Name, message)
	}
}

func (p *Player) applyKick(ball *Ball){
	rand.Seed(time.Now().UnixNano())
	angle := 2 * math.Pi * rand.Float64()

	// put the ball in a larger distance than player kick threshold
	// TODO put the ball NEAR the threshold so sometimes he might re-kick the ball
	ball.X = p.X + 1.1 * kickThreshold * math.Cos(angle)
	ball.Y = p.Y + 1.1 * kickThreshold * math.Sin(angle)

	v := 1 + rand.Float64() * 2
	ball.Vx = v * math.Cos(angle)
	ball.Vy = v * math.Sin(angle)
	ball.HolderID = p.ID
	ball.HolderTeam = p.TeamID
	ball.LastKick = time.Now()

	p.LastKick = time.Now()

	p.log(fmt.Sprintf("\nKick!!! (angle: %d degrees, velocity: %f)\n", int(180 * angle / math.Pi), v))
}

func getBallInputChannel() <- chan *Ball {
	rawInput := messaging.GetInputChannel(messaging.BallChannelName)
	res := make(chan *Ball)

	// Ball channel population, executed in function closure
	go func(){
		for val := range rawInput {
			bs := &Ball{}
			err := json.Unmarshal(val, bs)
			if err == nil {
				res <- bs
			}
		}
	}()
	return res
}

func getBallOutputChannel() chan <- *Ball {
	rawOutput := messaging.GetOutputChannel(messaging.BallChannelName)
	res := make(chan *Ball)

	// Ball channel population, executed in function closure
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





