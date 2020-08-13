package coursera

import (
	"errors"
	"strings"
	"text/template"
)

var coursesTmpl *template.Template
var onDemandLectureVideosTmpl *template.Template

var errInvalidRequest = errors.New("invalid request")

func init() {
	var err error
	if coursesTmpl, err = template.New("courses.v1").
		Parse("https://www.coursera.org/api/courses.v1?q=slug&slug={{.Slug}}&fields=domainTypes"); err != nil {
		panic(err)
	}

	if onDemandLectureVideosTmpl, err = template.New("onDemandLectureVideos.v1").
		Parse("https://www.coursera.org/api/onDemandLectureVideos.v1/{{.ItemId}}~{{.VideoId}}?fields=videoId,video&includes=video"); err != nil {
		panic(err)
	}
}

type coursesApiRequest struct {
	Slug string
}

func coursesApiLink(slug string) (string, error) {
	if slug == "" {
		return "", errInvalidRequest
	}

	s := new(strings.Builder)
	if err := coursesTmpl.Execute(s, coursesApiRequest{Slug: slug}); err != nil {
		return "", err
	}

	return s.String(), nil
}

type onDemandLectureVideoApiRequest struct {
	VideoId string
	ItemId  string
}

func onDemandLectureVideoApiLink(videoId, itemId string) (string, error) {
	if videoId == "" || itemId == "" {
		return "", errInvalidRequest
	}

	s := new(strings.Builder)
	if err := onDemandLectureVideosTmpl.Execute(s, onDemandLectureVideoApiRequest{
		VideoId: videoId,
		ItemId:  itemId,
	}); err != nil {
		return "", err
	}

	return s.String(), nil
}
