package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/miclle/lisa/action"
	"github.com/miclle/lisa/msg"
)

var version = "0.0.1-dev"

const usage = `A file watcher cli.

Usage: lisa COMMAND [ARGS]

All commands can be run with -h (or --help) for more information.

More info https://github.com/miclle/lisa
`

var authors = []cli.Author{
	cli.Author{Name: "Miclle", Email: "miclle.zheng@gmail.com"},
	cli.Author{Name: "Lisa", Email: "lisa_smiles@sina.com"},
}

func main() {
	app := cli.NewApp()
	app.Name = "lisa"
	app.Usage = usage
	app.Version = version
	app.Authors = authors

	app.CommandNotFound = func(c *cli.Context, command string) {
		msg.ExitCode(99)
		msg.Die("Command %s does not exist.", command)
	}

	app.Before = startup

	app.Commands = commands()

	if err := app.Run(os.Args); err != nil {
		msg.Err(err.Error())
		os.Exit(1)
	}

	// If there was a Error message exit non-zero.
	if msg.HasErrored() {
		m := msg.Color(msg.Red, "An Error has occurred")
		msg.Msg(m)
		os.Exit(2)
	}
}

func startup(c *cli.Context) error {
	// TODO
	return nil
}

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:        "server",
			ShortName:   "s",
			Usage:       "Serving Static Files with HTTP",
			Description: "Serving Static Files with HTTP",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Value: 8080,
					Usage: "Serving Static Files with HTTP used port.",
				},
				cli.StringFlag{
					Name:  "dir, d",
					Value: "./",
					Usage: "Serving Static Files with HTTP in directory.",
				},
				cli.StringFlag{
					Name:  "bind, b",
					Value: "0.0.0.0",
					Usage: "Serving Static Files with HTTP bind address.",
				},
			},
			Action: func(c *cli.Context) {
				action.Server(c.Int("port"), c.String("bind"), c.String("dir"))
			},
		},
		{
			Name:        "watch",
			ShortName:   "w",
			Usage:       "Start the lisa watcher",
			Description: "Start the lisa watcher",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir, d",
					Value: "./",
					Usage: "Watching the directory.",
				},
			},
			Action: func(c *cli.Context) {
				action.Watcher(c.String("dir"))
			},
		},
	}
}
