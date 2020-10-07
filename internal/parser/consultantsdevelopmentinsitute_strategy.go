package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

type ConsultantsdevelopmentinsituteStrategy struct{
	*commonStrategy
}

func NewConsultantsdevelopmentinsituteStrategy(log *logger.Logger, options *options) *ConsultantsdevelopmentinsituteStrategy {
	return &ConsultantsdevelopmentinsituteStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *ConsultantsdevelopmentinsituteStrategy) Name() string {
	return "consultantsdevelopmentinsitute.org"
}

func (s *ConsultantsdevelopmentinsituteStrategy) Run(project.ConfigRow) error {
	//return errors.New("has not been implemented yet")
	return nil
}
