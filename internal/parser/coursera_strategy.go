package parser

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/parser/coursera"
	"github.com/choizydev/LRP-tool/internal/project"
	"github.com/choizydev/LRP-tool/internal/storage"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type CourseraStrategy struct {
	*commonStrategy
	collectorInst *colly.Collector
	client        *coursera.Client
}

func NewCourseraStrategy(log *logger.Logger, options *options) *CourseraStrategy {
	return &CourseraStrategy{
		commonStrategy: newCommonStrategy(log, options),
		client:         coursera.NewClient(),
	}
}

func (s *CourseraStrategy) log() *logrus.Entry {
	return s.commonStrategy.log.FieldLogger().
		WithField("strategy", s.Name())
}

func (s *CourseraStrategy) Name() string {
	return "coursera.org"
}

func (s *CourseraStrategy) Run(configRow project.ConfigRow) error {
	c := s.collector()

	c.OnHTML("a[data-click-key=\"video_logged_out_page.video_lecture_landing_page.click.other_videos_link\"]", s.visitLinkHandler())
	c.OnHTML("a[data-e2e=\"course-link\"]", s.visitLinkHandler())
	c.OnHTML("a[data-click-key=\"xdp_v1.xdp.click.syllabus_item_link\"]", s.syllabusItemLinkHandler(configRow))

	if err := c.Visit(configRow.Link); err != nil {
		return err
	}
	c.Wait()

	return nil
}

func (s *CourseraStrategy) collector() *colly.Collector {
	if s.collectorInst != nil {
		return s.collectorInst.Clone()
	}

	collector := colly.NewCollector(
		colly.Async(s.options.async),
		colly.AllowedDomains("coursera.org", "www.coursera.org"),
		colly.CacheDir(s.options.cacheDir),
		colly.Debugger(&debug.LogDebugger{Output: s.log().WriterLevel(logrus.DebugLevel)}),
	)

	if err := collector.Limit(&colly.LimitRule{
		Delay:      time.Second,
		DomainGlob: "*coursera.org*",
	}); err != nil {
		s.log().WithError(err).Errorln("Can not set parser limits")
	}

	s.collectorInst = collector

	return collector
}

func (s *CourseraStrategy) syllabusItemLinkHandler(configRow project.ConfigRow) colly.HTMLCallback {
	return func(el *colly.HTMLElement) {
		itemUrl := el.Attr("href")
		splitUrl := strings.Split(itemUrl, "/")

		//Split link to get course slug and ID
		slug := splitUrl[2]
		s.log().WithField("slug", slug).Debugln("Fetching a course data from API")
		courses, err := s.client.Courses(slug)
		if err != nil {
			s.log().WithError(err).Errorln("Can not fetch the course data from API")
			return
		}

		first, err := courses.First()
		if err != nil {
			s.log().
				WithError(err).
				WithField("slug", slug).
				Errorln("Can not get the course data from API response")
			return
		}

		splitLastUrlSegment := strings.Split(splitUrl[3], "-")
		videoId := splitLastUrlSegment[len(splitLastUrlSegment)-1] // get last
		itemId := first.ID
		s.log().
			WithField("video_id", videoId).
			WithField("item_id", itemId).
			Debugln("Fetching a course video data from API")
		videos, err := s.client.Videos(videoId, itemId)
		if err != nil {
			s.log().
				WithError(err).
				WithField("video_id", videoId).
				WithField("item_id", itemId).
				Errorln("Can not get the course video data from API response")
			return
		}

		// store founded data using storage.Storager
		s.store(configRow.Scope, configRow.Profession, first.Name, el.Request.URL.String(), videos)

		// open page to parse other video link
		s.visitLink(el, itemUrl)
	}
}

func (s *CourseraStrategy) visitLinkHandler() colly.HTMLCallback {
	return func(el *colly.HTMLElement) {
		s.visitLink(el, el.Attr("href"))
	}
}

func (s *CourseraStrategy) visitLink(el *colly.HTMLElement, link string) {
	if err := el.Request.Visit(link); err != nil {
		if err != colly.ErrAlreadyVisited {
			s.log().WithField("link", link).WithError(err).Errorln("Link parse failed")
		}
	}
}

func (s *CourseraStrategy) store(
	scope string,
	profession string,
	name string,
	url string,
	videos *coursera.OnDemandLectureVideos,
) {
	if s.options.storage == nil {
		return
	}

	videosByResolution, err := videos.VideosByResolution()
	if err != nil {
		s.log().WithError(err).Warnln("video links are empty")
		return
	}

	if err := s.options.storage.Save(
		storage.Entity{
			Scope:      scope,
			Profession: profession,
			Title:      name,
			Link:       url,
			Movie:      videosByResolution.R720P.Mp4VideoURL,
		}); err != nil {
		s.log().WithError(err).Errorln("storage.Storager.Save")
	}
}
