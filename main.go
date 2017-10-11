package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	var sourceHost string
	var sourcePort int
	var targetHost string
	var targetPort int

	app := cli.NewApp()
	app.Name = "udpd"
	app.Version = "1.0.0"
	app.Usage = "Simple AF UDP proxy server"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "src_host",
			Value:       "127.0.0.1",
			Destination: &sourceHost,
			Usage:       "Host to bind",
		},
		cli.IntFlag{
			Name:        "src_port",
			Value:       12345,
			Destination: &sourcePort,
			Usage:       "Port to listen to",
		},
		cli.StringFlag{
			Name:        "dst_host",
			Value:       "127.0.0.1",
			Destination: &targetHost,
			Usage:       "Host to forward to",
		},
		cli.IntFlag{
			Name:        "dst_port",
			Value:       54321,
			Destination: &targetPort,
			Usage:       "Port to forward to",
		},
	}

	app.Action = func(c *cli.Context) error {
		runServer(sourceHost, sourcePort, targetHost, targetPort)
		return nil
	}

	app.Run(os.Args)
}
