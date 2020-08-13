package parser

import (
	"context"
	"errors"
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type UdemyStrategy struct {
	*commonStrategy
}

func NewUdemyStrategy(ctx context.Context, log *logger.Logger, options *options) *UdemyStrategy {
	return &UdemyStrategy{commonStrategy: newCommonStrategy(ctx, log, options)}
}

func (s *UdemyStrategy) Name() string {
	return "udemy.com"
}

func (s *UdemyStrategy) Run(_ project.ConfigRow) error {
	return errors.New("has not been implemented yet")
}
