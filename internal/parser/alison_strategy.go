package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type AlisonStrategy struct {
	*commonStrategy
}

func NewAlisonStrategy(log *logger.Logger, options *options) *AlisonStrategy {
	return &AlisonStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *AlisonStrategy) Name() string {
	return "alison.com"
}

func (s *AlisonStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
