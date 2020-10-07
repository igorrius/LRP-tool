package parser

import (
	"errors"
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/igorrius/LRP-tool/internal/project"
)

var ErrStrategyNotImplemented = errors.New("strategy has not been implemented yet")

type Strategy interface {
	Name() string
	Run(project.ConfigRow) error
}

func buildStrategy(log *logger.Logger, options *options, origin string) (Strategy, error) {
	switch origin {
	case "coursera.org", "https://www.coursera.org":
		return NewCourseraStrategy(log, options), nil
	case "www.udemy.com", "udemy.com", "www.udeymy.com", "https://www.udemy.com":
		return NewUdemyStrategy(log, options), nil
	case "futurelearn.com", "www.futurelearn.com", "https://www.futurelearn.com":
		return NewFutureLearnStrategy(log, options), nil
	case "edx.org", "https://www.edx.org":
		return NewFEdxStrategy(log, options), nil
	case "learndigital.withgoogle.com":
		return NewLearndigitalWithgoogleStrategy(log, options), nil
	case "udacity.com", "www.udacity.com", "https://www.udacity.com":
		return NewUdacityStrategy(log, options), nil
	case "plusacumen.org":
		return NewPlusacumenStrategy(log, options), nil
	case "philanthropyu.org":
		return NewPlusacumenStrategy(log, options), nil
	case "youtube.com", "www.youtube.com", "https://www.youtube.com":
		return NewYoutubeStrategy(log, options), nil
	case "consultantsdevelopmentinsitute.org":
		return NewConsultantsdevelopmentinsituteStrategy(log, options), nil
	case "alison.com":
		return NewAlisonStrategy(log, options), nil
	case "www.coacademy.com":
		return NewCoacademyStrategy(log, options), nil
	case "skillshare.com":
		return NewSkillshareStrategy(log, options), nil
	case "onlinehaircourses.com":
		return NewOnlinehaircoursesStrategy(log, options), nil
	case "istudy.org.uk":
		return NewIstudyStrategy(log, options), nil
	case "kadenze.com":
		return NewKadenzeStrategy(log, options), nil
	case "study.com":
		return NewStudyStrategy(log, options), nil
	case "ravensbourne.ac.uk":
		return NewRavensbourneStrategy(log, options), nil
	case "bolc.co.uk":
		return NewBolcStrategy(log, options), nil
	case "oxfordhomestudy.com":
		return NewOxfordhomestudyStrategy(log, options), nil
	case "www.khanacademy.org":
		return NewKhanacademyStrategy(log, options), nil
	case "canvas.net":
		return NewCanvasStrategy(log, options), nil
	case "journalismcourses.org":
		return NewJournalismcoursesStrategy(log, options), nil
	case "olympic.org":
		return NewOlympicStrategy(log, options), nil
	case "", "ability to work under pressure", "commercial awareness", "critical thinking", "decision making",
		"flexibility":
		return NewErrorStrategy(log, options), nil
	default:
		log.FieldLogger().WithField("origin", origin).Errorln("Strategy has not implemented yet")
		return nil, ErrStrategyNotImplemented
	}
}
