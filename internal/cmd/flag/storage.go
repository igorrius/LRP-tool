package flag

import "github.com/urfave/cli/v2"

const (
	PostgreSQLStorageName = "storage.postgresql"
	DummyStorageName      = "storage.dummy"
)

func DummyStorage() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    DummyStorageName,
		Usage:   "Enable Dummy engine database as a parsing result storage",
		EnvVars: []string{"APP_STORAGE_DUMMY"},
	}
}
func PostgreSQL() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    PostgreSQLStorageName,
		Usage:   "Enable PostgreSQL database as a parsing result storage",
		EnvVars: []string{"APP_STORAGE_POSTGRESQL_DSN"},
	}
}
