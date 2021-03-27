package main

import (
	"fmt"
	"log"
	"os"

	"rare/cmd"
	"rare/pkg/color"
	"rare/pkg/humanize"
	"rare/pkg/multiterm"
	"rare/pkg/multiterm/termunicode"

	"github.com/urfave/cli"
)

func cliMain(args ...string) error {
	app := cli.NewApp()

	app.Usage = "A fast regex parser, extractor and realtime aggregator"

	app.Version = fmt.Sprintf("%s, %s", version, buildSha)

	app.Description = `Aggregate and display information parsed from text files using
	regex and a simple handlebars-like expression syntax.

	Run "rare docs overview" for more information
	
	https://github.com/zix99/rare`

	app.Copyright = `rare  Copyright (C) 2019 Chris LaPointe
	This program comes with ABSOLUTELY NO WARRANTY.
	This is free software, and you are welcome to redistribute it
	under certain conditions`

	app.UseShortOptionHandling = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "nocolor,nc",
			Usage: "Disables color output",
		},
		cli.BoolFlag{
			Name:  "noformat,nf",
			Usage: "Disable number formatting",
		},
		cli.BoolFlag{
			Name:  "nounicode,nu",
			Usage: "Disable usage of unicode characters",
		},
		cli.BoolFlag{
			Name:  "color",
			Usage: "Force-enable color output",
		},
		cli.BoolFlag{
			Name:  "notrim",
			Usage: "By default, rare will trim output text for in-place updates. Setting this flag will disable that",
		},
		cli.StringFlag{
			Name:  "profile",
			Usage: "Write application profiling information as part of execution. Specify base-name",
		},
	}

	app.Commands = cmd.GetSupportedCommands()
	app.Commands = append(app.Commands, cli.Command{
		Name:   "_gendoc",
		Hidden: true,
		Usage:  "Generates documentation",
		Action: func(c *cli.Context) error {
			text, _ := c.App.ToMarkdown()
			fmt.Print(text)
			return nil
		},
	})

	app.Before = cli.BeforeFunc(func(c *cli.Context) error {
		if c.Bool("nocolor") {
			color.Enabled = false
		} else if c.Bool("color") {
			color.Enabled = true
		}
		if c.Bool("noformat") {
			humanize.Enabled = false
		}
		if c.Bool("notrim") {
			multiterm.AutoTrim = false
		}
		if c.Bool("nounicode") {
			termunicode.UnicodeEnabled = false
		}

		// Profiling
		if c.IsSet("profile") {
			basename := c.String("profile")
			startProfiler(basename)
		}

		return nil
	})

	app.After = cli.AfterFunc(func(c *cli.Context) error {
		if c.IsSet("profile") {
			stopProfile()
		}
		return nil
	})

	return app.Run(args)
}

func main() {
	err := cliMain(os.Args...)
	if err != nil {
		log.Fatal(err)
	}
}
