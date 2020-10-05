package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type EdxStrategy struct {
	*commonStrategy
}

func NewFEdxStrategy(log *logger.Logger, options *options) *EdxStrategy {
	return &EdxStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *EdxStrategy) Name() string {
	return "udemy.com"
}

func (s *EdxStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
