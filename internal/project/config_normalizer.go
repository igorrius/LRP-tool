package project

import (
	"errors"
	"strings"
)

var errEmptyConfOriginSite = errors.New("configuration origin site link is empty")

func normalizeConfigRow(r ConfigRow) (ConfigRow, error) {
	if r.Link == "" {
		return ConfigRow{}, errEmptyConfOriginSite
	}

	return ConfigRow{
		Types:          normalizeConfigType(r.Types),
		Specialization: normalizeConfigSpecialization(r.Specialization),
		OriginSite:     normalizeConfigOriginSite(r.OriginSite),
		Courses:        normalizeConfigCourses(r.Courses),
		Subscription:   normalizeConfigSubscription(r.Subscription),
		Link:           normalizeConfigLink(r.Link),
	}, nil
}

func normalizeConfigLink(s string) string {
	return strings.ToLower(s)
}

func normalizeConfigSubscription(s string) string {
	return strings.ToLower(s)
}

func normalizeConfigCourses(s string) string {
	return strings.ToLower(s)
}

func normalizeConfigSpecialization(s string) string {
	return strings.ToLower(s)
}

func normalizeConfigType(s string) string {
	return strings.ToLower(s)
}

func normalizeConfigOriginSite(s string) string {
	return strings.ToLower(s)
}
