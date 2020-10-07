package parser

import (
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/sirupsen/logrus"
	"time"
)

type UdemyStrategy struct {
	*commonStrategy
	collectorInst *colly.Collector
}

func NewUdemyStrategy(log *logger.Logger, options *options) *UdemyStrategy {
	return &UdemyStrategy{commonStrategy: newCommonStrategy(log, options)}
}

func (s *UdemyStrategy) log() *logrus.Entry {
	return s.commonStrategy.log.FieldLogger().WithField("strategy", s.Name())
}

func (s *UdemyStrategy) Name() string {
	return "udemy.com"
}

func (s *UdemyStrategy) Run(configRow project.ConfigRow) error {
	c := s.collector()

	//if err := c.Visit(configRow.Link); err != nil {
	//	return err
	//}
	c.Wait()

	return nil
}

func (s *UdemyStrategy) collector() *colly.Collector {
	if s.collectorInst != nil {
		return s.collectorInst.Clone()
	}

	collector := colly.NewCollector(
		colly.Async(s.options.async),
		colly.AllowedDomains("udemy.org", "www.udemy.org"),
		colly.CacheDir(s.options.cacheDir),
		colly.Debugger(&debug.LogDebugger{Output: s.log().WriterLevel(logrus.DebugLevel)}),
	)

	if err := collector.Limit(&colly.LimitRule{
		Delay: time.Second,
		DomainGlob: "*udemy.org*",
	}); err != nil {
		s.log().WithError(err).Errorln("Can not set parser limits")
	}

	s.collectorInst = collector

	return collector
}
