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

	app.Commands = []cli.Command{
		commands.JoinCommand,
		commands.ThrowCommand,
		commands.SimulateCommand,
	}

	app.Run(os.Args)


}


