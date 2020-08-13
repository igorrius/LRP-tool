package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigurations_Normalized(t *testing.T) {
	tests := []struct {
		name string
		c    Configurations
		want Configurations
	}{
		{
			name: "",
			c: Configurations{
				ConfigRow{
					Types:          "Types1",
					Specialization: "Specialization1",
					OriginSite:     "OriginSite1",
					Courses:        "Courses1",
					Subscription:   "Subscription1",
					Link:           "https://Link/route",
				},
				ConfigRow{
					Types:          "Types2",
					Specialization: "Specialization2",
					OriginSite:     "OriginSite2",
					Courses:        "Courses2",
					Subscription:   "Subscription2",
					Link:           "",
				},
				ConfigRow{
					Types:          "Types3",
					Specialization: "Specialization3",
					OriginSite:     "OriginSite3",
					Courses:        "Courses3",
					Subscription:   "Subscription3",
					Link:           "https://Link3/route",
				},
			},
			want: Configurations{
				ConfigRow{
					Types:          "types1",
					Specialization: "specialization1",
					OriginSite:     "originsite1",
					Courses:        "courses1",
					Subscription:   "subscription1",
					Link:           "https://link/route",
				},
				ConfigRow{
					Types:          "types3",
					Specialization: "specialization3",
					OriginSite:     "originsite3",
					Courses:        "courses3",
					Subscription:   "subscription3",
					Link:           "https://link3/route",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.c.Normalized())
		})
	}
}
