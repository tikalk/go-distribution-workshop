package commands

import (
	"github.com/urfave/cli"
	"github.com/tikalk/go-distribution-workshop/models"
	"fmt"
	"github.com/tikalk/go-distribution-workshop/messaging"
	"sync"
	"github.com/tikalk/go-distribution-workshop/apps"
	"strings"
)

var JoinCommand = cli.Command{
	Name: 	"join",
	Usage:  "Join an existing game",
	Action: joinGame,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "players",
			Usage: "comma-separated list of player names",
			Value: "Roy",
		},
		cli.StringFlag{
			Name:  "team",
			Usage: fmt.Sprintf("team to assign players to " +
					"(%s / %s / both). On 'both' players will be assigned equally to both team",
					string(models.TeamBlue),
					string(models.TeamRed),
				),
			Value: string(models.TeamBlue),
		},
	},
}

func joinGame(c *cli.Context) error {
	defer messaging.Stop()
	setupRedis(c)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	playersFlag := c.String("players")
	players := strings.Split(playersFlag, ",")

	teamFlag := c.String("team")

	go apps.JoinGame(players, models.Team(teamFlag), wg)

	wg.Wait()

	return nil
}
