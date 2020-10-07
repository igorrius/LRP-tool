package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type FutureLearnStrategy struct {
	*commonStrategy
}

func NewFutureLearnStrategy(log *logger.Logger, options *options) *FutureLearnStrategy {
	return &FutureLearnStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *FutureLearnStrategy) Name() string {
	return "udemy.com"
}

func (s *FutureLearnStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
