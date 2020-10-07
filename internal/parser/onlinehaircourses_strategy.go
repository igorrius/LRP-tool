package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type OnlinehaircoursesStrategy struct {
	*commonStrategy
}

func NewOnlinehaircoursesStrategy(log *logger.Logger, options *options) *OnlinehaircoursesStrategy {
	return &OnlinehaircoursesStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *OnlinehaircoursesStrategy) Name() string {
	return "onlinehaircourses.com"
}

func (s *OnlinehaircoursesStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
