package commands

import (
	"github.com/urfave/cli"
	"github.com/tikalk/go-distribution-workshop/models"
	"fmt"
)

var PlayCommand = cli.Command{
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
			Usage: fmt.Sprintf("team to assign players to " +
					"(%s / %s / both). On 'both' players will be assigned equally to both team",
					string(models.TeamBlue),
					string(models.TeamRed),
				),
			Value: string(models.TeamBlue),
		},
	},
}

func assignPlayers(c *cli.Context) error {
	return nil
}
