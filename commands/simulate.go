package commands

import (
	"sync"
	"github.com/tikalk/go-distribution-workshop/apps"
	"github.com/urfave/cli"
	"time"
	"fmt"
	"github.com/mgutz/ansi"
)

var SimulateCommand = cli.Command{
		Name:  "simulate",
		Usage: "Run an End-to-End game simulation",
		Action: simulate,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "players",
				Usage: "total number of players - distributed evenly to teams",
				Value: 4,
			},
			cli.IntFlag{
				Name:  "port",
				Usage: "port of display server",
				Value: 8080,
			},
		},
	}

func simulate(c *cli.Context) error {
	setupRedis(c)

	wg := &sync.WaitGroup{}
	wg.Add(2)


	port := c.Int("port")
	go apps.LaunchDisplay(port, wg)

	players := c.Int("players")
	go apps.ExecuteSimulation(players, wg)

	time.Sleep(200 * time.Millisecond)
	fmt.Printf(ansi.Color("\n\nDisplay server launched successfully on port %d\n", "green"), port)

	wg.Wait()

	return nil
}

