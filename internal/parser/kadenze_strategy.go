package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type KadenzeStrategy struct {
	*commonStrategy
}

func NewKadenzeStrategy(log *logger.Logger, options *options) *KadenzeStrategy {
	return &KadenzeStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *KadenzeStrategy) Name() string {
	return "kadenze.com"
}

func (s *KadenzeStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
