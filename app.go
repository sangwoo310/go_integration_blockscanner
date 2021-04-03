package main

import (
	"fmt"
	hf "github.com/sangwoo310/go_framework_hsw"
	"github.com/urfave/cli/v2"
)

// set process control
func main()  {
	//app := &cli.App{
	//	Commands:Commands2,
	//}
	//
	//app.Run(os.Args)
	s := hf.SHswFrame{Name:"hsw"}
	hf.NewFrame(s)

}

var Commands2 []*cli.Command

func init() {
	Commands2 = []*cli.Command{
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
			Subcommands:            []*cli.Command{

			},
			Flags:                  []cli.Flag{
				&cli.StringFlag{
					Name:        "rpc",
					Usage:       "rpc test",
					Required:    false,
				},

			},
			HelpName:               "helpping test : ",
		},
	}
}

func test(ctx *cli.Context) error  {
	fmt.Println("test commands ")

	fmt.Println(ctx.String("rpc"))

	return nil
}