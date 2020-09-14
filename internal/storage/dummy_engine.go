package storage

import (
	"context"
	"github.com/choizydev/LRP-tool/internal/logger"
)

type DummyEngine struct {
	log   *logger.Logger
	count int
}

func NewDummyEngine(ctx context.Context) *DummyEngine {
	return &DummyEngine{log: logger.FromContextOrNew(ctx)}
}

func (d *DummyEngine) Save(entity Entity) error {
	d.count++
	d.log.FieldLogger().
		WithField("count", d.count).
		WithField("entity", entity).
		Debugln("DummyEngine.Save")

	return nil
}

func (d *DummyEngine) String() string {
	return "Dummy engine"
}

func (d *DummyEngine) Close() error {
	return nil
}
