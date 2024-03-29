package commands

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/urfave/cli"
	"go-distribution-fuzeday/apps"
	"sync"
	"time"
)

var DisplayCommand = cli.Command{
	Name:   "display",
	Usage:  "Launch display server for an existing game",
	Action: display,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "port",
			Usage: "port of display server",
			Value: 8080,
		},
	},
}

func display(c *cli.Context) error {

	setupRedis(c)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	port := c.Int("port")
	go apps.LaunchDisplay(port, wg)

	time.Sleep(100 * time.Millisecond)
	fmt.Printf(ansi.Color("Display server launched successfully on port %d\n", "green"), port)

	wg.Wait()
	return nil
}
