package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type IstudyStrategy struct {
	*commonStrategy
}

func NewIstudyStrategy(log *logger.Logger, options *options) *IstudyStrategy {
	return &IstudyStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *IstudyStrategy) Name() string {
	return "istudy.org.uk"
}

func (s *IstudyStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
