package project

import (
	"encoding/csv"
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/jszwec/csvutil"
	"os"
	"path/filepath"
	"strings"
)

func normalizeDirPath(dirPath string) string {
	return strings.TrimSuffix(dirPath, "/")
}

func loadConfigurations(confDir string, log *logger.Logger) (Configurations, error) {
	log.FieldLogger().WithField("conf_dir", confDir).Debugln("Call loadConfigurations()")
	conf := make(Configurations, 0, 64)

	err := filepath.Walk(normalizeDirPath(confDir),
		func(path string, info os.FileInfo, err error) error {
			if !strings.HasSuffix(info.Name(), ".csv") {
				return nil
			}

			fileNameWithDir := confDir + "/" + info.Name()
			log.FieldLogger().WithField("config_file_name", info.Name()).Infoln("Processing config file")
			f, err := os.Open(fileNameWithDir)
			if err != nil {
				return err
			}

			decoder, err := csvutil.NewDecoder(csv.NewReader(f))
			if err != nil {
				return err
			}

			var c Configurations
			if err := decoder.Decode(&c); err != nil {
				return err
			}

			conf = append(conf, c...)

			return err
		})
	if err != nil {
		return nil, err
	}

	return conf, nil
}
