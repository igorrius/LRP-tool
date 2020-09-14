package coursera

import (
	"errors"
)

var ErrFirstCourseNotFound = errors.New("the first course has not been found")
var ErrVideoNotFound = errors.New("video not found")

type CoursesRow struct {
	CourseType  string `json:"courseType"`
	DomainTypes []struct {
		SubdomainID string `json:"subdomainId"`
		DomainID    string `json:"domainId"`
	} `json:"domainTypes"`
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Courses struct {
	Elements []CoursesRow `json:"elements"`
	Paging   struct {
		Total int `json:"total"`
	} `json:"paging"`
	Linked struct {
	} `json:"linked"`
}

func (c *Courses) First() (*CoursesRow, error) {
	if len(c.Elements) > 0 {
		return &c.Elements[0], nil
	}
	return nil, ErrFirstCourseNotFound
}

type ByResolutionVideo struct {
	R720P struct {
		WebMVideoURL    string `json:"webMVideoUrl"`
		PreviewImageURL string `json:"previewImageUrl"`
		Mp4VideoURL     string `json:"mp4VideoUrl"`
	} `json:"720p"`
	R360P struct {
		WebMVideoURL    string `json:"webMVideoUrl"`
		PreviewImageURL string `json:"previewImageUrl"`
		Mp4VideoURL     string `json:"mp4VideoUrl"`
	} `json:"360p"`
	R540P struct {
		WebMVideoURL    string `json:"webMVideoUrl"`
		PreviewImageURL string `json:"previewImageUrl"`
		Mp4VideoURL     string `json:"mp4VideoUrl"`
	} `json:"540p"`
}

type OnDemandLectureVideos struct {
	Elements []struct {
		ItemID   string `json:"itemId"`
		VideoID  string `json:"videoId"`
		ID       string `json:"id"`
		CourseID string `json:"courseId"`
	} `json:"elements"`
	Paging struct {
	} `json:"paging"`
	Linked struct {
		OnDemandVideosV1 []struct {
			Subtitles struct {
				Ar   string `json:"ar"`
				Ru   string `json:"ru"`
				ZhTW string `json:"zh-TW"`
				Uk   string `json:"uk"`
				PtBR string `json:"pt-BR"`
				El   string `json:"el"`
				En   string `json:"en"`
				ID   string `json:"id"`
				ZhCN string `json:"zh-CN"`
				Tr   string `json:"tr"`
				Es   string `json:"es"`
			} `json:"subtitles"`
			Sources struct {
				Playlists struct {
					Hls string `json:"hls"`
				} `json:"playlists"`
				ByResolution ByResolutionVideo `json:"byResolution"`
			} `json:"sources"`
			SubtitlesVtt struct {
				Ar   string `json:"ar"`
				Ru   string `json:"ru"`
				ZhTW string `json:"zh-TW"`
				Uk   string `json:"uk"`
				PtBR string `json:"pt-BR"`
				El   string `json:"el"`
				En   string `json:"en"`
				ID   string `json:"id"`
				ZhCN string `json:"zh-CN"`
				Tr   string `json:"tr"`
				Es   string `json:"es"`
			} `json:"subtitlesVtt"`
			ID string `json:"id"`
		} `json:"onDemandVideos.v1"`
	} `json:"linked"`
}

func (v *OnDemandLectureVideos) VideosByResolution() (ByResolutionVideo, error) {
	for _, s := range v.Linked.OnDemandVideosV1 {
		return s.Sources.ByResolution, nil
	}
	return ByResolutionVideo{}, ErrVideoNotFound
}
