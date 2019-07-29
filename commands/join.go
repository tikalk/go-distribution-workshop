package commands

import (
	"github.com/urfave/cli"
	"github.com/tikalk/go-distribution-workshop/models"
	"fmt"
	"sync"
	"github.com/tikalk/go-distribution-workshop/apps"
	"strings"
	"errors"
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
					string(models.Brazil),
					string(models.Argentina),
				),
			Value: string(models.Brazil),
		},
	},
}

func joinGame(c *cli.Context) error {

	playersFlag := c.String("players")
	players := strings.Split(playersFlag, ",")

	teamFlag := c.String("team")


	if len(players) == 0 {
		println("Al least one player must be specified")
		return errors.New("Al least one player must be specified")
	}

	switch teamFlag {
	case string(models.Brazil):
	case string(models.Argentina):
	case string(models.Both):
		break
	default:
		errMsg := fmt.Sprintf("Illegal value  for --team flag. Must be one of {%s, %s, both}", models.Brazil, models.Argentina)
		println(errMsg)
		return errors.New(errMsg)
	}


	wg := &sync.WaitGroup{}
	wg.Add(1)







	go apps.JoinGame(players, models.Team(teamFlag), wg)

	wg.Wait()

	return nil
}
