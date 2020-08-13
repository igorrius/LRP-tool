package flag

import "github.com/urfave/cli/v2"

const (
	JsonOutputFormatterName = "json"
	LogLevelName = "log-level"
)

func LogLevel() cli.Flag {
	return &cli.StringFlag{
		Name:        LogLevelName,
		Usage:       "Set logger level to one of [panic, fatal, error, warn, warning, info, debug, trace]",
		DefaultText: "info",
		Value:       "info",
		EnvVars:     []string{"APP_LOG_LEVEL"},
		Aliases:     []string{"ll"},
	}
}

func JsonOutputFormatter() cli.Flag {
	return &cli.BoolFlag{
		Name:        JsonOutputFormatterName,
		Usage:       "Use json as output format",
		Value:       false,
		DefaultText: "false",
		EnvVars:     []string{"APP_LOG_FORMATTER_JSON"},
	}
}
