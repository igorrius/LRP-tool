package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type PlusacumenStrategy struct {
	*commonStrategy
}

func NewPlusacumenStrategy(log *logger.Logger, options *options) *PlusacumenStrategy {
	return &PlusacumenStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *PlusacumenStrategy) Name() string {
	return "plusacumen.org"
}

func (s *PlusacumenStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
