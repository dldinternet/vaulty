package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/vaulty/vaulty"
)

var config = vaulty.NewConfig()

var proxyCommand = &cli.Command{
	Name:  "proxy",
	Usage: "run proxy server",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "enable debug (exposes request and response bodies)",
			Destination: &config.Debug,
		},
		&cli.StringFlag{
			Name:        "address",
			Aliases:     []string{"a"},
			Value:       ":8080",
			Usage:       "address that vaulty should listen on",
			Destination: &config.Address,
		},
		&cli.StringFlag{
			Name:        "routes-file",
			Aliases:     []string{"r"},
			Value:       "./routes.json",
			Usage:       "routes file",
			Destination: &config.RoutesFile,
		},
		&cli.StringFlag{
			Name:        "ca-path",
			Aliases:     []string{"ca"},
			Value:       "./",
			Usage:       "path to CA key and cert",
			Destination: &config.CAPath,
		},
		&cli.StringFlag{
			Name:        "proxy-pass",
			Aliases:     []string{"p"},
			Usage:       "forward proxy password",
			EnvVars:     []string{"PROXY_PASS"},
			Destination: &config.ProxyPassword,
		},
		&cli.StringFlag{
			Name:        "key",
			Aliases:     []string{"k"},
			Usage:       "forward proxy password",
			EnvVars:     []string{"ENCRYPTION_KEY"},
			Destination: &config.EncryptionKey,
		},
	},
	Action: func(c *cli.Context) error {
		if err := config.GenerateMissedValues(); err != nil {
			return fmt.Errorf("Error with generating missed values: %s", err)
		}

		return vaulty.Run(config)
	},
}
