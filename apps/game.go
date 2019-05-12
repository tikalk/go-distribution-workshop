package apps

import (
	"time"
	"encoding/json"
	"math/rand"
	"sync"
	"fmt"
	"github.com/tikalk/go-distribution-workshop/messaging"
	"github.com/tikalk/go-distribution-workshop/models"
	"github.com/satori/go.uuid"
)

func JoinGame(players []string, team models.Team, externalWaitGroup *sync.WaitGroup) {

	displayChannel := getDisplayOutputChannel()
	rand.Seed(time.Now().UnixNano())

	numPlayers := len(players)
	wg := &sync.WaitGroup{}
	wg.Add(numPlayers)


	for i := 0; i < numPlayers; i++ {
		u2 := uuid.Must(uuid.NewV4())

		player := &models.Player{
			ID: u2.String(),
			Name: players[i],
			X: rand.Float64() * 100,
			Y: rand.Float64() * 100,
			MaxVelocity: rand.Float64() * 0.1,
		}
		if (team == models.Both && i %2 == 0) || team == models.Argentina {
			player.TeamID = models.Argentina
		} else{
			player.TeamID = models.Brazil
		}
		fmt.Printf("Added player %s\n", players[i])
		player.Activate(displayChannel, wg)

	}

	wg.Wait()

	if externalWaitGroup != nil {
		externalWaitGroup.Done()
	}
}

func ExecuteSimulation(numPlayers int, externalWaitGroup *sync.WaitGroup) {

	ThrowBall(-1, -1)


	displayChannel := getDisplayOutputChannel()
	rand.Seed(time.Now().UnixNano())
	wg := &sync.WaitGroup{}
	wg.Add(numPlayers)

	fmt.Println("Adding players...")

	for i := 0; i < numPlayers; i++ {

		u2 := uuid.Must(uuid.NewV4())

		player := &models.Player{
			ID: u2.String(),
			X: rand.Float64() * 100,
			Y: rand.Float64() * 100,
			MaxVelocity: rand.Float64() * 0.1,
		}
		if i % 2 == 0 {
			player.TeamID = models.Argentina
		} else {
			player.TeamID = models.Brazil
		}

		player.Name = models.GetPlayerName(player.TeamID)
		player.Activate(displayChannel, wg)
	}

	wg.Wait()
	if externalWaitGroup != nil {
		externalWaitGroup.Done()
	}

}

func getDisplayOutputChannel() chan <- *models.DisplayStatus  {
	rawOutput := messaging.GetOutputChannel(messaging.DisplayChannelName)
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

func ThrowBall(x, y float64){
	rand.Seed(time.Now().UnixNano())

	if x == -1 {
		x = rand.Float64() * 100.0
	}
	if y == -1 {
		y = rand.Float64() * 100.0
	}

	fmt.Println("Throwing ball!")

	bs := &models.Ball{X: x, Y: y, Vx: 0, Vy: 0, Vz: 0, Z: 50}
	bs.LastUpdated = time.Now()

	models.GetBallOutputChannel() <- bs

}
