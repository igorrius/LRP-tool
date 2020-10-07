package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type BolcStrategy struct {
	*commonStrategy
}

func NewBolcStrategy(log *logger.Logger, options *options) *BolcStrategy {
	return &BolcStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func(*BolcStrategy) Name() string {
	return "bolc.co.uk"
}

func(*BolcStrategy) Run(configRow project.ConfigRow) error {
	return nil
}