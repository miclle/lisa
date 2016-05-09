package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/miclle/lisa/action"
	"github.com/miclle/lisa/msg"
	"github.com/skratchdot/open-golang/open"
)

var version = "0.1.0"

const usage = `A file watcher cli.

Usage: lisa COMMAND [ARGS]

All commands can be run with -h (or --help) for more information.

More info http://miclle.me/lisa/
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
			Usage:       "Starting a file system watcher",
			Description: "Starting a file system watcher then execute a command",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "command, c",
					Value: "",
					Usage: "Execute the command when the directory files modified.",
				},
				cli.StringFlag{
					Name:  "path, p",
					Value: "./",
					Usage: "Watching the directory or file.",
				},
				cli.StringFlag{
					Name:  "event, e",
					Value: "create,rename,write,remove",
					Usage: "Execute the command when the events was trigger: create,rename,write,remove,chmod",
				},
				cli.IntFlag{
					Name:  "delay, d",
					Value: 1000,
					Usage: "Execute the command after the number of milliseconds.",
				},
			},
			Action: func(c *cli.Context) {
				action.Watcher(c.String("path"), c.String("event"), c.String("command"), c.Int("delay"))
			},
		},
		{
			Name:  "home",
			Usage: "Go to the document website",
			Action: func(c *cli.Context) {
				open.Run("http://miclle.me/lisa/")
			},
		},
	}
}
