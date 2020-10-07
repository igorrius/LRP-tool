package parser

import (
	"errors"
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

var ErrStrategyNotImplemented = errors.New("strategy has not been implemented yet")

type Strategy interface {
	Name() string
	Run(project.ConfigRow) error
}

func buildStrategy(log *logger.Logger, options *options, origin string) (Strategy, error) {
	switch origin {
	case "coursera.org", "https://www.coursera.org":
		return NewCourseraStrategy(log, options), nil
	default:
		return nil, ErrStrategyNotImplemented
	}
}
