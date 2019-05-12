package models

import (
	"math/rand"
	"math"
	"time"
	"fmt"
)

type (
	Team string
)

const (
	Brazil    Team = "brazil"
	Argentina Team = "argentina"
	Both      Team = "both"
)

var PlayersBrazil = []string{"Gérson", "Tostão", "Zico", "Rivaldo", "Jairzinho", "Ronaldo", "Carlos Alberto", "Didi", "Garrincha", "Pelé"}
var PlayersArgentina = []string{"Crespo", "Messi", "Ricardo", "Carrizo", "Labruna", "Batistuta", "Ubaldo", "Kempes", "Passarella", "Maradona"}

func init(){
	rand.Seed(time.Now().UnixNano())
	randomisePlayers(PlayersBrazil)
	randomisePlayers(PlayersArgentina)
}
var indexBrazil = 0
var indexArgentina = 0
var totalPlayerRequests = 0;

func GetPlayerName(team Team) string{
	res := fmt.Sprintf("Player %d", totalPlayerRequests)

	switch team {
	case Brazil:
		if indexBrazil < len(PlayersBrazil){
			res = PlayersBrazil[indexBrazil]
		}
		indexBrazil++
	case Argentina:
		if indexArgentina < len(PlayersArgentina) {
			res = PlayersArgentina[indexArgentina]
		}
		indexArgentina++
	}
	totalPlayerRequests++

	return res;
}
func randomisePlayers(players []string){
	for i := range players {
		other := i + int(math.Floor(rand.Float64() * float64(len(players) - i)))
		players[i], players[other] = players[other], players[i]
	}
}