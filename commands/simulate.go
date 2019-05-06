package commands

import (
	"sync"
	"github.com/tikalk/go-distribution-workshop/apps"
	"github.com/urfave/cli"
)

var SimulateCommand = cli.Command{
		Name:  "simulate",
		Usage: "Run an End-to-End game simulation",
		Action: simulate,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "players",
				Usage: "total number of players - distributed evenly to teams",
				Value: 1,
			},
			cli.IntFlag{
				Name:  "port",
				Usage: "port of display server",
				Value: 8080,
			},
		},
	}

func simulate(c *cli.Context) error {

	wg := &sync.WaitGroup{}
	wg.Add(1)


	players := c.Int("players")

	go apps.ExecuteSimulation(players, wg)

	wg.Wait()

	return nil
}

