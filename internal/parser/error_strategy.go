package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type ErrorStrategy struct {
	*commonStrategy
}

func NewErrorStrategy(log *logger.Logger, options *options) *ErrorStrategy {
	return &ErrorStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *ErrorStrategy) Name() string {
	return ""
}

func (s *ErrorStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
