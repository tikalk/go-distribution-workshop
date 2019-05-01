package main

import (
	"github.com/urfave/cli"
	"flag"
	"os"
	"github.com/tikalk/go-distribution-workshop/commands"
)
var app cli.App

func init(){
	app = cli.App{
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
		commands.PlayCommand,
		commands.ThrowCommand,
		commands.SimulateCommand,
		commands.DisplayCommand,

	}


}
func main()  {



	flag.Parse()
	app.Run(os.Args)
}


