package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type OxfordhomestudyStrategy struct {
	*commonStrategy
}

func NewOxfordhomestudyStrategy(log *logger.Logger, options *options) *OxfordhomestudyStrategy {
	return &OxfordhomestudyStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *OxfordhomestudyStrategy) Name() string {
	return "oxfordhomestudy.com"
}

func (s *OxfordhomestudyStrategy) Run(configRow project.ConfigRow) error{
	return nil
}
