package models

import (
	"time"
	"math"
	"github.com/tikalk/go-distribution-workshop/utils"
)

type (
	Ball struct {
		X float64					`json:"x"`
		Y float64					`json:"y"`
		Vx float64					`json:"v_x"`
		Vy float64					`json:"v_y"`
		Z float64					`json:"z"`
		Vz float64					`json:"vz"`
		LastPlayer string			`json:"last_player"`
		LastTeam	int				`json:"last_team"`
		LastUpdated	time.Time		`json:"last_updated"`
		HolderID string				`json:"holder_id"`
		HolderTeam Team				`json:"holder_team"`
		LastKick time.Time			`json:"last_kick"`
	}
)


const EnergyLoss = 0.96
const GlobalDumping = 0.98
const g = 0.098


func (b *Ball) GetSurfaceVelocity() float64{
	return math.Sqrt(math.Pow(b.Vx, 2) + math.Pow(b.Vy, 2))
}

func (b *Ball) ApplyKinematics(){
	timeDiff := time.Now().Sub(b.LastUpdated).Seconds()
	iterations := int(math.Max(timeDiff / 0.01, 1))

	for i := 0; i < iterations; i++ {
		b.applyKinematicsIteration(timeDiff, float64(iterations))
	}
}
func (b *Ball) applyKinematicsIteration(timeDiff, iterations float64){
	effectiveDumping := GlobalDumping //1 - ((1 - GlobalDumping) / iterations)
	effectiveG := g / iterations


	b.Vx *= effectiveDumping
	b.Vy *= effectiveDumping
	b.Vz -= effectiveG
	b.Vz *= effectiveDumping

	utils.ApplyVelocityComponent(&b.X, &b.Vx, 1.0, iterations)
	utils.ApplyVelocityComponent(&b.Y, &b.Vy, 1.0, iterations)
	utils.ApplyVelocityComponent(&b.Z, &b.Vz, EnergyLoss, iterations)

}

var ballChannel chan *Ball

func GetBallChannel() chan *Ball {
	if ballChannel == nil {
		ballChannel = make(chan *Ball, 1024)
	}

	return ballChannel
}





