package project

import (
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/sirupsen/logrus"
)

type ConfigRow struct {
	Courses        string `json:"courses" csv:"courses"`
	Link           string `json:"link" csv:"link"`
	OriginSite     string `json:"origin_site" csv:"origin_site"`
	Specialization string `json:"specialization" csv:"specialization"`
	Subscription   string `json:"subscription" csv:"subscription"`
	Types          string `json:"types" csv:"types"`
}

func (r *ConfigRow) ToLogrusFields() logrus.Fields {
	return logrus.Fields{
		"courses":        r.Courses,
		"link":           r.Link,
		"origin_site":    r.OriginSite,
		"specialization": r.Specialization,
		"subscription":   r.Subscription,
		"types":          r.Types,
	}
}

type Configurations []ConfigRow

func (c Configurations) Normalized() Configurations {
	normalizedConf := make(Configurations, 0, len(c))
	for _, configRow := range c {
		row, err := normalizeConfigRow(configRow)
		if err != nil {
			continue
		}
		normalizedConf = append(normalizedConf, row)
	}
	return normalizedConf
}

//Config contains logic that manage project configuration files and provide it for other components
type Config struct {
	configurations Configurations
}

//NewConfig
// confDir - parameter that sets folder to find configuration files
func NewConfig(confDir string, log *logger.Logger) *Config {
	c, err := loadConfigurations(confDir, log)
	if err != nil {
		log.FieldLogger().WithError(err).Fatalln("Can not load configuration")
	}
	return &Config{configurations: c}
}

func (c *Config) Configurations() Configurations {
	return c.configurations
}
