package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd = &cli.App{}
}

func setCommands()  {
	Commands2 := []*cli.Command{
		{
			Name:                   "test",
			Aliases:                nil,
			Usage:                  "test command",
			UsageText:              "test command for testing",
			Description:            "description",
			ArgsUsage:              "No args",
			Category:               "wow",
			Action:                 test,
			OnUsageError:           nil,
			Subcommands:            nil,
			Flags:                  nil,
			HelpName:               "test help",
		},
	}
	cmd.Commands = Commands2

}

func sample() {
	Commands = []cli.Command{
		{
			Name:                   "test",
			Aliases:                nil,
			Usage:                  "test command",
			UsageText:              "test command for testing",
			Description:            "description",
			ArgsUsage:              "No args",
			Category:               "wow",
			Action:                 test,
			OnUsageError:           nil,
			Subcommands:            nil,
			Flags:                  nil,
			HelpName:               "test help",
		},
	}
}

func test(ctx *cli.Context) error  {
	fmt.Println("test commands ")

	return nil
}