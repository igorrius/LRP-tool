package storage

import (
	"context"
	"errors"
	"github.com/choizydev/LRP-tool/internal/logger"
)

var (
	ErrFailedStorageChannel = errors.New("sending to the storage async channel has been failed")
)

type (
	VideoResolution = string
	VideoSourceUrl  = string
)

type Entity struct {
	Name   string
	Link   string
	Movies map[VideoResolution]VideoSourceUrl
}

type Storager interface {
	Save(Entity) error
}

type Storage struct {
	ctx      context.Context
	logger   *logger.Logger
	engines  []Engine
	entities chan *Entity
}

func NewStorage(ctx context.Context, options ...Option) *Storage {
	log := logger.FromContextOrNew(ctx)

	s := &Storage{
		ctx:      ctx,
		logger:   log,
		engines:  []Engine{newDummyEngine(log)},
		entities: make(chan *Entity, 8),
	}

	// Loop through each option
	for _, option := range options {
		// Call the option giving the instantiated
		// *Storage as the argument
		option(s)
	}

	go s.startStorageHandler()

	return s
}

func (s *Storage) Save(entity Entity) error {
	select {
	case s.entities <- &entity:
		return nil
	default:
		if err := s.ctx.Err(); err != nil {
			return err
		}

		return ErrFailedStorageChannel
	}
}

func (s *Storage) startStorageHandler() {
	go func() {
		<-s.ctx.Done()
		close(s.entities)
	}()

	for entity := range s.entities {
		for _, engine := range s.engines {
			if err := engine.Save(*entity); err != nil {
				s.logger.FieldLogger().WithError(err).WithField("engine", engine).Errorln("Engine.Save")
			}
		}
	}
}
