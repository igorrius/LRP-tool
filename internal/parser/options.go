package parser

import "github.com/igorrius/LRP-tool/internal/storage"

type Option func(options *options)

type options struct {
	async    bool
	cacheDir string
	storage  storage.Storager
}

func newOptions(opts ...Option) *options {
	// create options with default values
	o := &options{
		async:    false,
		cacheDir: "./.cache",
	}

	// apply options that set by arguments
	for _, opt := range opts {
		opt(o)
	}

	return o
}

//Async option enable async pipeline in the parser useless load
//default is false
func Async(b bool) Option {
	return func(options *options) {
		options.async = b
	}
}

//Set the folder path to enable the sites parsed cache to prevent useless load
//default is ./.cache
//supported by the [coursera] strategies
func CacheDir(dir string) Option {
	return func(options *options) {
		options.cacheDir = dir
	}
}

//enable storage
func Storage(storage storage.Storager) Option {
	return func(options *options) {
		options.storage = storage
	}
}
