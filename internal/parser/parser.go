package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type Parser struct {
	conf project.Configurations
	log  *logger.Logger
}

func NewParser(log *logger.Logger, c project.Configurations) *Parser {
	return &Parser{conf: c, log: log}
}

func (p *Parser) Run(opts ...Option) error {
	options := newOptions(opts...)

	for _, configRow := range p.conf {
		strategy, err := buildStrategy(p.log, options, configRow.OriginSite)
		if err != nil {
			return err
		}

		p.log.FieldLogger().WithField("strategy", strategy.Name()).Infoln("Trying to start parsing strategy")
		if err := strategy.Run(configRow); err != nil {
			p.log.FieldLogger().
				WithField("strategy", strategy.Name()).
				WithFields(configRow.ToLogrusFields()).
				WithError(err).Errorln("Parse error")
		}
	}

	return nil
}
