package main

import (
	"github.com/urfave/cli"
	"go-distribution-fuzeday/commands"
	"os"
)

func main() {

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
		commands.DisplayCommand,
	}

	app.Run(os.Args)
}
