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
	defer messaging.Stop()

	displayChannel := getDisplayOutputChannel()
	rand.Seed(time.Now().UnixNano())

	numPlayers := len(players)
	wg := sync.WaitGroup{}
	wg.Add(numPlayers)


	for i := 0; i < numPlayers; i++ {
		u2 := uuid.NewV4()

		player := &models.Player{
			ID: u2.String(),
			Name: players[i],
			X: rand.Float64() * 100,
			Y: rand.Float64() * 100,
			MaxVelocity: rand.Float64() * 0.1,
		}
		if (team == models.TeamBoth && i %2 == 0) || team == models.TeamRed{
			player.TeamID = models.TeamRed
		} else{
			player.TeamID = models.TeamBlue
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
	defer messaging.Stop()

	throwBall()


	displayChannel := getDisplayOutputChannel()
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	wg.Add(numPlayers)

	fmt.Println("Adding players...")

	for i := 0; i < numPlayers; i++ {

		u2 := uuid.NewV4()

		player := &models.Player{
			ID: u2.String(),
			Name: fmt.Sprintf("Player %d", i),		// TODO get names from list
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
	if externalWaitGroup != nil {
		externalWaitGroup.Done()
	}

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
	fmt.Println("Throwing ball!")
	output, _ := messaging.GetOutputChannel(messaging.BallChannelName)

	bs := &models.Ball{X: 8, Y: 2, Vx: 0, Vy: 0, Vz: 0, Z: 50}
	bs.LastUpdated = time.Now()
	bsSer, err := json.Marshal(bs)
	if err != nil {
		panic(err)
	}
	output <- bsSer
}
