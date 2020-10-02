package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type JournalismcoursesStrategy struct {
	*commonStrategy
}

func NewJournalismcoursesStrategy(log *logger.Logger, options *options) *JournalismcoursesStrategy {
	return &JournalismcoursesStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *JournalismcoursesStrategy) Name() string {
	return "journalismcourses.org"
}

func (s *JournalismcoursesStrategy) Run(configRow project.ConfigRow) error {
	return nil
}
