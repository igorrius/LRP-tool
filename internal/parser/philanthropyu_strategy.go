package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type PhilanthropyuStrategy struct {
	*commonStrategy
}

func NewPhilanthropyuStrategy(log *logger.Logger, options *options) *PhilanthropyuStrategy {
	return &PhilanthropyuStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *PhilanthropyuStrategy) Name() string {
	return "philanthropyu.org"
}

func (s *PhilanthropyuStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
