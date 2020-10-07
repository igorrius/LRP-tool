package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type SkillshareStrategy struct {
	*commonStrategy
}

func NewSkillshareStrategy(log *logger.Logger, options *options) *SkillshareStrategy {
	return &SkillshareStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *SkillshareStrategy) Name() string {
	return "www.coacademy.com"
}

func (s *SkillshareStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}