package storage

type Option func(storage *Storage)

func WithEngine(engine Engine) Option {
	return func(storage *Storage) {
		storage.engines = append(storage.engines, engine)
	}
}
