package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type StudyStrategy struct {
	*commonStrategy
}

func NewStudyStrategy(log *logger.Logger, options *options) *StudyStrategy {
	return &StudyStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *StudyStrategy) Name() string {
	return "study.com"
}

func (s *StudyStrategy) Run(configRow project.ConfigRow) error {
	return nil
}
