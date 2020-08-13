package coursera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByResolutionVideo_toMap(t *testing.T) {
	type fields struct {
		R720P struct {
			WebMVideoURL    string `json:"webMVideoUrl"`
			PreviewImageURL string `json:"previewImageUrl"`
			Mp4VideoURL     string `json:"mp4VideoUrl"`
		}
		R360P struct {
			WebMVideoURL    string `json:"webMVideoUrl"`
			PreviewImageURL string `json:"previewImageUrl"`
			Mp4VideoURL     string `json:"mp4VideoUrl"`
		}
		R540P struct {
			WebMVideoURL    string `json:"webMVideoUrl"`
			PreviewImageURL string `json:"previewImageUrl"`
			Mp4VideoURL     string `json:"mp4VideoUrl"`
		}
	}
	tests := []struct {
		name   string
		fields fields
		want   map[VideoResolution]VideoSourceUrl
	}{
		{
			name: "All resolutions",
			fields: fields{
				R720P: struct {
					WebMVideoURL    string `json:"webMVideoUrl"`
					PreviewImageURL string `json:"previewImageUrl"`
					Mp4VideoURL     string `json:"mp4VideoUrl"`
				}{Mp4VideoURL: "R720P - Mp4VideoURL"},
				R360P: struct {
					WebMVideoURL    string `json:"webMVideoUrl"`
					PreviewImageURL string `json:"previewImageUrl"`
					Mp4VideoURL     string `json:"mp4VideoUrl"`
				}{Mp4VideoURL: "R360O - Mp4VideoURL"},
				R540P: struct {
					WebMVideoURL    string `json:"webMVideoUrl"`
					PreviewImageURL string `json:"previewImageUrl"`
					Mp4VideoURL     string `json:"mp4VideoUrl"`
				}{Mp4VideoURL: "R540P - Mp4VideoURL"},
			},
			want: map[string]string{
				"360p": "R360O - Mp4VideoURL",
				"540p": "R540P - Mp4VideoURL",
				"720p": "R720P - Mp4VideoURL"},
		},
		{
			name: "Only 360p",
			fields: fields{
				R360P: struct {
					WebMVideoURL    string `json:"webMVideoUrl"`
					PreviewImageURL string `json:"previewImageUrl"`
					Mp4VideoURL     string `json:"mp4VideoUrl"`
				}{Mp4VideoURL: "360P - Mp4VideoURL"},
			},
			want: map[string]string{"360p": "360P - Mp4VideoURL"},
		},
		{
			name: "Only 540p",
			fields: fields{
				R540P: struct {
					WebMVideoURL    string `json:"webMVideoUrl"`
					PreviewImageURL string `json:"previewImageUrl"`
					Mp4VideoURL     string `json:"mp4VideoUrl"`
				}{Mp4VideoURL: "R540P - Mp4VideoURL"},
			},
			want: map[string]string{"540p": "R540P - Mp4VideoURL"},
		},
		{
			name: "Only 720p",
			fields: fields{
				R720P: struct {
					WebMVideoURL    string `json:"webMVideoUrl"`
					PreviewImageURL string `json:"previewImageUrl"`
					Mp4VideoURL     string `json:"mp4VideoUrl"`
				}{Mp4VideoURL: "R720P - Mp4VideoURL"},
			},
			want: map[string]string{"720p": "R720P - Mp4VideoURL"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ByResolutionVideo{
				R720P: tt.fields.R720P,
				R360P: tt.fields.R360P,
				R540P: tt.fields.R540P,
			}

			assert.Equal(t, tt.want, v.ToMap())
		})
	}
}
