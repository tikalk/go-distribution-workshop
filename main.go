package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/tikalk/go-distribution-workshop/commands"
)

// TODO add dependencies to both Project.dep, sh and bat build scripts

func main()  {

	app := cli.NewApp()

	app.Name = "go-distribution-workshop"
	app.Version = "1.0.0"
	app.Email = "royp@tikalk.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "redis-host",
			Usage: "IP of Redis server",
			Value: "127.0.0.1",
		},
		cli.IntFlag{
			Name:  "redis-port",
			Usage: "port of Redis server",
			Value: 6379,
		},
	}

	app.Commands = []cli.Command{
		commands.JoinCommand,
		commands.ThrowCommand,
		commands.SimulateCommand,
	}

	app.Run(os.Args)


}


