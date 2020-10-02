package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
)

type KhanacademyStrategy struct {
	*commonStrategy
}

func NewKhanacademyStrategy(log *logger.Logger, options *options) *KhanacademyStrategy {
	return &KhanacademyStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s KhanacademyStrategy) Name() string {
	return "www.khanacademy.org"
}

func (s KhanacademyStrategy) Run(configRow project.ConfigRow) error {
	return nil
}
