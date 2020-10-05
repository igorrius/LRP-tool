package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type RavensbourneStrategy struct {
	*commonStrategy
}

func NewRavensbourneStrategy(log *logger.Logger, options *options) *RavensbourneStrategy {
	return &RavensbourneStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *RavensbourneStrategy) Name() string {
	return "ravensbourne.ac.uk"
}

func (s *RavensbourneStrategy) Run(configRow project.ConfigRow) error {
	return nil
}
