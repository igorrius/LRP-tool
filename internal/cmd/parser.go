package cmd

import (
	"github.com/igorrius/LRP-tool/internal/cmd/parser"
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/urfave/cli/v2"
)

func Parser() *cli.Command {
	return &cli.Command{
		Name:  "parser",
		Usage: "Parser command section",
		Before: func(ctx *cli.Context) error {
			log := logger.FromContextOrNew(ctx.Context)
			log.FieldLogger().WithField("cmd_name", ctx.Args().First()).Debugln("Running command")
			return nil
		},
		Subcommands: []*cli.Command{
			parser.Run(),
		},
	}
}
