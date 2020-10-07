package cmd

import (
	"github.com/igorrius/LRP-tool/internal/cmd/config"
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/urfave/cli/v2"
)

func Config() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Config tools",
		Before: func(ctx *cli.Context) error {
			log := logger.FromContextOrNew(ctx.Context)
			log.FieldLogger().WithField("cmd_name", ctx.Args().First()).Debugln("Running command")
			return nil
		},
		Subcommands: []*cli.Command{
			config.Import(),
			config.Inspect(),
		},
	}
}
