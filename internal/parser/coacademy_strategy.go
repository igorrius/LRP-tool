package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type CoacademyStrategy struct {
	*commonStrategy
}

func NewCoacademyStrategy(log *logger.Logger, options *options) *CoacademyStrategy {
	return &CoacademyStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *CoacademyStrategy) Name() string {
	return "www.coacademy.com"
}

func (s *CoacademyStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
