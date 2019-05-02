package commands

import "github.com/urfave/cli"

var ThrowCommand = cli.Command{
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
}

func throwBall(c *cli.Context) error {
	setupRedis(c)

	return nil
}
