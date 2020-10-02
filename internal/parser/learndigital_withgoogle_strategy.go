package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type LearndigitalWithgoogleStrategy struct {
	*commonStrategy
}

func NewLearndigitalWithgoogleStrategy(log *logger.Logger, options *options) *LearndigitalWithgoogleStrategy {
	return &LearndigitalWithgoogleStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *LearndigitalWithgoogleStrategy) Name() string {
	return "learndigital.withgoogle.com"
}

func (s *LearndigitalWithgoogleStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}

