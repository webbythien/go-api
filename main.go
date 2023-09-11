package main

import (
	"fmt"
	"os"

	"log"

	"lido-core/v1/cmd"
	"lido-core/v1/pkg/utils"

	"github.com/urfave/cli"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var (
	client *cli.App
)

func init() {
	// Initialise a CLI app
	client = cli.NewApp()
	client.Name = "lido core api"
	client.Usage = "lido core api worker and handler"
	client.Version = "0.0.1"
}

func main() {

	client.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				log.Printf("start %s\n", c.Args().First())
				consume := fmt.Sprintf("core-api-consume:%s", utils.NewID())
				if err := cmd.WorkerExecute(c.Args().First(), consume, 12); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "server",
			Usage: "start api server",
			Action: func(c *cli.Context) error {
				log.Printf("start %s\n", c.Command.Name)
				cmd.StartServer()
				return nil
			},
		},
	}

	// Run the CLI app
	client.Run(os.Args)
}
