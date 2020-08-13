package storage

import "github.com/choizydev/LRP-tool/internal/logger"

type dummyEngine struct {
	log   *logger.Logger
	count int
}

func newDummyEngine(log *logger.Logger) *dummyEngine {
	return &dummyEngine{log: log}
}

func (d *dummyEngine) Save(entity Entity) error {
	d.count++
	d.log.FieldLogger().
		WithField("count", d.count).
		WithField("entity", entity).
		Debugln("dummyEngine.Save")

	return nil
}

func (d *dummyEngine) String() string {
	return "Dummy engine"
}
