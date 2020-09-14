package config

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/project"
	"github.com/jszwec/csvutil"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func Import() *cli.Command {
	return &cli.Command{
		Name:  "import",
		Usage: "Import file based configuration from different sources",
		Action: func(ctx *cli.Context) error {
			log := logger.FromContextOrNew(ctx.Context)
			log.FieldLogger().Infoln("Import configuration")

			srcFolder := "./business/categories"
			dstFolder := "./config/project"

			return importProjectConfigs(srcFolder, dstFolder, log)
		},
	}
}

var supportedFiles = []string{".xlsx", ".xls"}

func importProjectConfigs(srcFolder string, dstFolder string, log *logger.Logger) error {
	if err := filepath.Walk(
		srcFolder,
		func(path string, info os.FileInfo, err error) error {
			if !isFileSupported(info.Name(), supportedFiles...) {
				return nil
			}

			fileNameWithDir := srcFolder + "/" + info.Name()
			scope := extractScope(info.Name())
			log.FieldLogger().
				WithField("scope", scope).
				WithField("source_file_name", fileNameWithDir).
				Infoln("Processing source file")

			f, err := excelize.OpenFile(fileNameWithDir)
			if err != nil {
				fmt.Println(err)
				return err
			}

			sheetName := f.GetSheetName(0)
			log.FieldLogger().WithField("sheet_name", strings.TrimSpace(sheetName)).Debugln("Sheet name")

			rows, err := f.GetRows(sheetName)
			if err != nil {
				return err
			}

			configRows := make([]project.ConfigRow, 0, 64)
			var lastSpecialization string
			for i, col := range rows {
				if i < 2 || len(col) < 7 {
					continue
				}

				// if specialization column
				if specialization := col[2]; specialization != "" {
					lastSpecialization = specialization
					log.FieldLogger().WithField("specialization", specialization).Debugln("Process specialization")
					continue
				}

				configRows = append(configRows, project.ConfigRow{
					Courses:      col[4],
					Link:         col[6],
					OriginSite:   normalizeOriginSite(col[3]),
					Profession:   lastSpecialization,
					Subscription: col[5],
					Scope:        scope,
				})
			}

			// write a data to a destination file
			destinationFileName := fmt.Sprintf("%s/%s.csv", dstFolder, scope)
			file, err := os.Create(destinationFileName)
			if err != nil {
				return err
			}
			defer func() { _ = file.Close() }()

			log.FieldLogger().
				WithField("source_file_name", fileNameWithDir).
				WithField("destination_file_name", destinationFileName).
				Infoln("Save a project config data")

			writer := csv.NewWriter(file)
			defer writer.Flush()
			encoder := csvutil.NewEncoder(writer)

			return encoder.Encode(configRows)
		}); err != nil {
		return err
	}

	return nil
}

func normalizeOriginSite(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		if r == '/' || r == ' ' {
			return true
		}
		return false
	})
}

func extractScope(fileName string) string {
	return strings.TrimLeftFunc(
		strings.TrimSuffix(fileName, filepath.Ext(fileName)),
		func(r rune) bool {
			if !unicode.IsLetter(r) {
				return true
			}
			return false
		},
	)
}

func isFileSupported(fileName string, extensions ...string) bool {
	for _, extension := range extensions {
		if strings.HasSuffix(fileName, extension) {
			return true
		}
	}

	return false
}
