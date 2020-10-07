package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type UdacityStrategy struct {
	*commonStrategy
}

func NewUdacityStrategy(log *logger.Logger, options *options) *UdacityStrategy {
	return &UdacityStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *UdacityStrategy) Name() string {
	return "udacity.com"
}

func (s *UdacityStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}