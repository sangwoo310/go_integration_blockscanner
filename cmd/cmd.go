package cmd

import (
	"github.com/urfave/cli/v2"
)

var cmd *cli.App
var Commands []cli.Command

func init() {
	cmd = &cli.App{}
}
