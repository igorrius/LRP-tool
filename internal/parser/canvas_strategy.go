package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type CanvasStrategy struct {
	*commonStrategy
}

func NewCanvasStrategy(log *logger.Logger, options *options) *CanvasStrategy {
	return &CanvasStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *CanvasStrategy) Name() string {
	return "canvas.net"
}

func (s *CanvasStrategy) Run(configRow project.ConfigRow) error {
	return nil
}
