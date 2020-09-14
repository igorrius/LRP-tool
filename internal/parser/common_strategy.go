package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
)

type commonStrategy struct {
	log     *logger.Logger
	options *options
}

func newCommonStrategy(log *logger.Logger, options *options) *commonStrategy {
	return &commonStrategy{log: log, options: options}
}
