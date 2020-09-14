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
		Scope:        normalizeConfigScope(r.Scope),
		Profession:   normalizeConfigProfession(r.Profession),
		OriginSite:   normalizeConfigOriginSite(r.OriginSite),
		Courses:      normalizeConfigCourses(r.Courses),
		Subscription: normalizeConfigSubscription(r.Subscription),
		Link:         normalizeConfigLink(r.Link),
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

func normalizeConfigProfession(s string) string {
	return s
}

func normalizeConfigScope(s string) string {
	return s
}

func normalizeConfigOriginSite(s string) string {
	return strings.ToLower(s)
}
