package parser

import (
	"context"
	"errors"
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

var ErrStrategyNotImplemented = errors.New("strategy has not been implemented yet")

type Strategy interface {
	Name() string
	Run(row project.ConfigRow) error
}

func buildStrategy(ctx context.Context, log *logger.Logger, options *options, origin string) (Strategy, error) {
	switch origin {
	case "coursera.org":
		return NewCourseraStrategy(ctx, log, options), nil
	case "udemy.com":
		return NewUdemyStrategy(ctx, log, options), nil
	default:
		return nil, ErrStrategyNotImplemented
	}
}
