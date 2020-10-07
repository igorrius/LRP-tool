package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type OlympicStrategy struct {
	*commonStrategy
}

func NewOlympicStrategy(log *logger.Logger, options *options) *OlympicStrategy {
	return &OlympicStrategy{ commonStrategy: newCommonStrategy(log, options)}
}

func (s *OlympicStrategy) Name() string {
	return "olympic.org"
}

func (s *OlympicStrategy) Run(configRow project.ConfigRow) error {
	return nil
}
