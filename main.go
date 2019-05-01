package main

import (
	"github.com/urfave/cli"
	"flag"
	"os"
)

func main()  {

	app := cli.App{
		Name:     "go-distribution-workshop",
		Version:  "1.0.0",
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "redis-host",
			Value: "127.0.0.1",
		},
		cli.StringFlag{
			Name:  "redis-port",
			Value: "6379",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "play",
			Usage: "Participate in an existing game",
			Action: assignPlayers,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "players",
					Usage: "comma-separated list of player names",
					Value: "Roy",
				},
				cli.StringFlag{
					Name:  "team",
					Usage: "team to assign players to (red / blue / both). On 'both' players will be assigned equally to both team",
					Value: "blue",
				},
			},
		},
		{
			Name:  "throw",
			Usage: "Throw a new ball to an existing game",
			Action: throwBall,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "x",
					Usage: "initial X coordinate of thrown ball",
					Value: 8,
				},
				cli.IntFlag{
					Name:  "y",
					Usage: "initial Y coordinate of thrown ball",
					Value: 5,
				},
				cli.BoolFlag{
					Name:  "manual_pos",
					Usage: "use specific coordinates to place the ball. Otherwise a random position will be used",
				},
			},
		},
		{
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
		},
		{
			Name:  "display",
			Usage: "Launch display server for an existing game",
			Action: display,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port",
					Usage: "port of display server",
					Value: 8080,
				},
			},
		},
	}


	flag.Parse()
	app.Run(os.Args)
}

func assignPlayers(c *cli.Context) error {
	return nil
}

func throwBall(c *cli.Context) error {
	return nil
}

func simulate(c *cli.Context) error {
	return nil
}

func display(c *cli.Context) error {
	return nil
}
