package parser

import (
	"errors"
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
	"sync"
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

	var wg sync.WaitGroup

	for _, configRow := range p.conf {
		strategy, err := buildStrategy(p.log, options, configRow.OriginSite)
		if err != nil {
			if errors.Is(err, ErrStrategyNotImplemented) {
				p.log.FieldLogger().
					WithField("origin", configRow.OriginSite).
					Errorln("Strategy has not implemented yet")
				continue
			}
			return err
		}

		wg.Add(1)
		go func(group *sync.WaitGroup, cf project.ConfigRow) {
			defer group.Done()
			p.log.FieldLogger().WithField("strategy", strategy.Name()).Infoln("Trying to start parsing strategy")
			if err := strategy.Run(cf); err != nil {
				p.log.FieldLogger().
					WithField("strategy", strategy.Name()).
					WithFields(cf.ToLogrusFields()).
					WithError(err).Errorln("Parse error")
			}
		}(&wg, configRow)
	}

	wg.Wait()

	return nil
}
