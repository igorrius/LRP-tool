package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type YoutubeStrategy struct{
	*commonStrategy
}

func NewYoutubeStrategy(log *logger.Logger, options *options) *YoutubeStrategy {
	return &YoutubeStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *YoutubeStrategy) Name() string {
	return "youtube.com"
}

func (s *YoutubeStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
