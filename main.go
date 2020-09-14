package main

import (
	"github.com/choizydev/LRP-tool/internal/cmd"
	"github.com/choizydev/LRP-tool/internal/cmd/flag"
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "LPR tool",
		Description: "Learning resources parser tool.",
		Version:     "v1.0.0",
		Usage:       "Ready for a new challenge in Goland with love!",
		Before: func(ctx *cli.Context) error {
			ctx.Context = logger.NewContext(
				ctx.Context,
				logger.NewLogger(createLoggerOptions(ctx)...))
			return nil
		},
		Commands: []*cli.Command{
			cmd.Config(),
			cmd.Parser(),
		},
		Flags: []cli.Flag{
			flag.LogLevel(),
			flag.JsonOutputFormatter(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.WithError(err).Fatal("Application start failed")
	}
}

func createLoggerOptions(ctx *cli.Context) []logger.Option {
	var opts []logger.Option

	// option that set logger output formatter to JSON
	if ctx.Bool(flag.JsonOutputFormatterName) {
		opts = append(opts, logger.WithJsonFormatter(true))
	}

	// option that provide logger level
	opts = append(opts, logger.WithLevel(ctx.String(flag.LogLevelName)))

	return opts
}
