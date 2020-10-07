package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/igorrius/LRP-tool/internal/logger"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/sirupsen/logrus"
)

const (
	stepsTable       = "steps"
	professionsTable = "professions"
	scopesTable      = "scopes"
)

type PostgresqlEngine struct {
	ctx    context.Context
	conn   *pgx.Conn
	logger *logger.Logger
}

func NewPostgresqlEngine(ctx context.Context, dsn string) *PostgresqlEngine {
	l := logger.FromContextOrNew(ctx)

	if config, err := pgx.ParseConfig(dsn); err == nil {
		// set application logger as a PGX logger
		config.Logger = logrusadapter.NewLogger(l.Logger())
		// try to set PGX logger level same with the application logger level
		if level, err := pgx.LogLevelFromString(l.Logger().Level.String()); err == nil {
			config.LogLevel = level
		}
	}

	// try to connect to the DB engine
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		l.FieldLogger().
			WithError(err).WithField("dsn", dsn).Fatalln("Unable to connect to database")
	}

	p := &PostgresqlEngine{ctx: ctx, logger: l, conn: conn}

	return p
}

func (e *PostgresqlEngine) log() *logrus.Entry {
	return e.logger.FieldLogger().
		WithField("storage", e.String())
}

func (e *PostgresqlEngine) Close() error {
	return e.conn.Close(e.ctx)
}

func (e *PostgresqlEngine) String() string {
	return "PostgreSQL engine"
}

func (e *PostgresqlEngine) Save(entity Entity) (saveErr error) {
	e.log().WithField("entity", entity).Debugln("Save")

	// 1 step - get profession by entity.professionsTable
	// 		if !present -> create a new profession row using scope id by entity.Scope (:title) and profession name
	// 2 step - save step using parser data and profession.ID

	// begin transaction
	tx, saveErr := e.conn.Begin(e.ctx)
	if saveErr != nil {
		return
	}
	defer func() {
		if saveErr != nil {
			return
		}

		if err := tx.Rollback(e.ctx); !errors.Is(err, pgx.ErrTxClosed) {
			saveErr = err
		}
	}()

	// get profession by entity.professionsTable
	professionId, saveErr := e.findOrCreateProfessionId(tx, entity)
	if saveErr != nil {
		return
	}

	// TODO: add steps saving logic
	_, saveErr = tx.Exec(
		e.ctx,
		fmt.Sprintf("INSERT INTO %s (title, profession_id, video_url, index) VALUES($1,$2,$3,1)", stepsTable),
		entity.Title, professionId, entity.Movie,
	)
	if saveErr != nil {
		return
	}

	return tx.Commit(e.ctx)
}

func (e *PostgresqlEngine) findOrCreateProfessionId(tx pgx.Tx, entity Entity) (int, error) {
	var professionId int
	if err := tx.QueryRow(
		e.ctx,
		fmt.Sprintf("SELECT id FROM %s WHERE name = $1 LIMIT 1;", professionsTable),
		entity.Profession,
	).Scan(&professionId); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return 0, err
	}

	// if profession has been founded than return profession ID
	if professionId > 0 {
		return professionId, nil
	}

	// create a new profession
	scopeId, err := e.findScopeId(tx, entity)
	if err != nil {
		return 0, err
	}

	if err := tx.QueryRow(
		e.ctx,
		fmt.Sprintf("INSERT INTO %s (name, scope_id) VALUES($1, $2) RETURNING id", professionsTable),
		entity.Profession, scopeId,
	).Scan(&professionId); err != nil {
		return 0, err
	}

	return professionId, nil
}

func (e *PostgresqlEngine) findScopeId(tx pgx.Tx, entity Entity) (id int, err error) {
	if err = tx.QueryRow(
		e.ctx,
		fmt.Sprintf("SELECT id FROM %s WHERE title = $1", scopesTable),
		entity.Scope,
	).Scan(&id); err != nil {
		return
	}

	return
}
