package parser

import (
	"errors"
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

var ErrStrategyNotImplemented = errors.New("strategy has not been implemented yet")

type Strategy interface {
	Name() string
	Run(project.ConfigRow) error
}

func buildStrategy(log *logger.Logger, options *options, origin string) (Strategy, error) {
	switch origin {
	case "coursera.org":
		return NewCourseraStrategy(log, options), nil
	case "www.udemy.com":
		return NewUdemyStrategy(log, options), nil
	case "futurelearn.com":
		return NewFutureLearnStrategy(log, options), nil
	case "edx.org":
		return NewFEdxStrategy(log, options), nil
	default:
		log.FieldLogger().WithField("origin", origin).Errorln("Strategy has not implemented yet")
		return nil, ErrStrategyNotImplemented
	}
}
