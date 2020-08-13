package parser

import (
	"context"
	"github.com/choizydev/LRP-tool/internal/logger"
)

type commonStrategy struct {
	ctx     context.Context
	log     *logger.Logger
	options *options
}

func newCommonStrategy(ctx context.Context, log *logger.Logger, options *options) *commonStrategy {
	return &commonStrategy{ctx: ctx, log: log, options: options}
}
