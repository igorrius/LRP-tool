package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
	"github.com/urfave/cli/v2"
)

func Inspect() *cli.Command {
	return &cli.Command{
		Name:  "inspect",
		Usage: "The project configuration inspection. This command shows the parser task configurations.",
		Action: func(ctx *cli.Context) error {
			log := logger.FromContextOrNew(ctx.Context)
			conf := project.NewConfig("./config/project", log)

			for _, configRow := range conf.Configurations().Normalized() {
				log.FieldLogger().WithFields(map[string]interface{}{
					"Types":          configRow.Types,
					"Specialization": configRow.Specialization,
					"OriginSite":     configRow.OriginSite,
					"Courses":        configRow.Courses,
					"Subscription":   configRow.Subscription,
					"Link":           configRow.Link,
				}).Infoln("Project config row")
			}

			return nil
		},
	}
}
