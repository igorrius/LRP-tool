package coursera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coursesLink(t *testing.T) {
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Empty slug",
			args:    args{slug: ""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "With valid argument",
			args:    args{slug: "valid-slug"},
			want:    "https://www.coursera.org/api/courses.v1?q=slug&slug=valid-slug&fields=domainTypes",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := coursesApiLink(tt.args.slug)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_onDemandLectureVideoApiLink(t *testing.T) {
	type args struct {
		videoId string
		itemId  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Invalid argument VideoId",
			args: args{
				videoId: "",
				itemId:  "itemID",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Invalid argument ItemId",
			args: args{
				videoId: "videoID",
				itemId:  "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Invalid both argument",
			args: args{
				videoId: "",
				itemId:  "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "With valid arguments",
			args: args{
				videoId: "video3030",
				itemId:  "item4040",
			},
			want:    "https://www.coursera.org/api/onDemandLectureVideos.v1/item4040~video3030?fields=videoId,video&includes=video",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := onDemandLectureVideoApiLink(tt.args.videoId, tt.args.itemId)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
