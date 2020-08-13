package coursera

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Courses(slug string) (*Courses, error) {
	requestUrl, err := coursesApiLink(slug)
	if err != nil {
		return nil, err
	}

	get, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	defer func() { _ = get.Body.Close() }()

	if get.StatusCode != http.StatusOK {
		return nil, errors.New("status code expected 200 but " + get.Status)
	}

	var res Courses
	decoder := json.NewDecoder(get.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Videos(videoId, itemId string) (*OnDemandLectureVideos, error) {
	requestUrl, err := onDemandLectureVideoApiLink(videoId, itemId)
	if err != nil {
		return nil, err
	}

	get, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	defer func() { _ = get.Body.Close() }()

	if get.StatusCode != http.StatusOK {
		return nil, errors.New("status code expected 200 but " + get.Status)
	}

	var res OnDemandLectureVideos
	decoder := json.NewDecoder(get.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
